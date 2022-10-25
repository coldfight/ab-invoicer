package invoice_generator

import (
	"encoding/json"
	"github.com/coldfight/ab-invoicer/internal/tools/pdf_generator"
	templateHelpers "github.com/coldfight/ab-invoicer/internal/tools/template_helpers"
	"html/template"
	"time"
)

const (
	TaxRate         = 0.13
	SavedDateLayout = "Jan 02, 2006"
)

type Date time.Time

func (d *Date) UnmarshalJSON(bytes []byte) error {
	var v interface{}
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}

	t, err := time.Parse(SavedDateLayout, v.(string))
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(SavedDateLayout))
}

func (d Date) Format(layout string) string {
	return time.Time(d).Format(layout)
}

func (d *Date) SetFromString(layout, dateStr string) error {
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return err
	}
	date := Date(t)
	*d = date
	return nil
}

type Expense struct {
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unitPrice"`
}
type ExpenseList []Expense

func (e Expense) TotalCost() float64 {
	return float64(e.Quantity) * e.UnitPrice
}

type Labour struct {
	Date        Date    `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}
type LabourList []Labour

func (l Labour) TotalCost() float64 {
	return l.Amount
}

func (el ExpenseList) ExpensesSubtotal() float64 {
	sum := 0.0
	for _, e := range el {
		sum += e.TotalCost()
	}
	return sum
}

func (el ExpenseList) ExpensesTaxes() float64 {
	return el.ExpensesSubtotal() * TaxRate
}

// @todo: Refactor this so we're caching the results and sending them to the template rather than calling the functions in the template
func (el ExpenseList) ExpensesWithTaxesSubtotal() float64 {
	return el.ExpensesSubtotal() + el.ExpensesTaxes()
}

func (ll LabourList) LabourSubtotal() float64 {
	sum := 0.0
	for _, l := range ll {
		sum += l.TotalCost()
	}
	return sum
}

func InvoiceTotal(el ExpenseList, ll LabourList) float64 {
	return el.ExpensesWithTaxesSubtotal() + ll.LabourSubtotal()
}

type Owner struct {
	Name       string `json:"name"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postalCode"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type BilledTo struct {
	Name       string `json:"name"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postalCode"`
	Phone      string `json:"phone"`
}

type Invoice struct {
	Owner         Owner       `json:"owner"`
	BilledTo      BilledTo    `json:"billedTo"`
	ExpenseList   ExpenseList `json:"expenseList"`
	LabourList    LabourList  `json:"labourList"`
	InvoiceNumber int         `json:"invoiceNumber"`
	InvoiceDate   Date        `json:"invoiceDate"`
}

type InvoiceTemplateData struct {
	Invoice
	InvoiceTotal        float64
	BootstrapStylesheet template.CSS
	Fonts               map[string]templateHelpers.FontVariation
}

func NewInvoice(invoice Invoice) {
	invoiceTotal := InvoiceTotal(invoice.ExpenseList, invoice.LabourList)
	bootstrapStylesheet := templateHelpers.GetStylesheet("assets/styles/bootstrap.css")
	fontMap := map[string]templateHelpers.FontVariation{
		"Normal": {
			Name:    "fira-code",
			Regular: templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-regular.ttf"),
			Bold:    templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-bold.ttf"),
		},
		"Mono": {
			Name:    "fira-code-mono",
			Regular: templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-regular-mono.ttf"),
			Bold:    templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-bold-mono.ttf"),
		},
	}

	templateData := InvoiceTemplateData{
		Invoice:             invoice,
		InvoiceTotal:        invoiceTotal,
		BootstrapStylesheet: bootstrapStylesheet,
		Fonts:               fontMap,
	}

	pdf_generator.CreatePdf("invoice.tmpl", "./storage/invoice.pdf", templateData)
}
