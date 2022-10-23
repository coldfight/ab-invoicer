package receipt_creator

import (
	"encoding/json"
	"github.com/coldfight/ab-invoicer/internal/tools"
	"io"
	"log"
	"os"
)

const (
	TaxRate = 0.13
)

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
	Date        tools.Date `json:"date"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount"`
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
	InvoiceDate   tools.Date  `json:"invoiceDate"`
}

func Create() {
	// Temp read from db.json file
	jsonFile, err := os.Open("./storage/db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var invoice Invoice
	err = json.Unmarshal(byteValue, &invoice)
	if err != nil {
		log.Fatal(err)
	}

	templateData := struct {
		ExpenseList   ExpenseList
		LabourList    LabourList
		BilledTo      BilledTo
		Owner         Owner
		InvoiceNumber int
		InvoiceDate   tools.Date
		InvoiceTotal  float64
	}{
		ExpenseList:   invoice.ExpenseList,
		LabourList:    invoice.LabourList,
		BilledTo:      invoice.BilledTo,
		Owner:         invoice.Owner,
		InvoiceNumber: invoice.InvoiceNumber,
		InvoiceDate:   invoice.InvoiceDate,
		InvoiceTotal:  InvoiceTotal(invoice.ExpenseList, invoice.LabourList),
	}

	tools.CreatePdf("invoice.tmpl", "./invoice.pdf", templateData)
}
