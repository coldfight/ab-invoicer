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

func (e Expense) TotalCost() float64 {
	return float64(e.Quantity) * e.UnitPrice
}

type Labour struct {
	Date        tools.Date `json:"date"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount"`
}

func (l Labour) TotalCost() float64 {
	return l.Amount
}

func ExpensesSubtotal(es []Expense) float64 {
	sum := 0.0
	for _, e := range es {
		sum += e.TotalCost()
	}
	return sum
}

func ExpensesTaxes(es []Expense) float64 {
	return ExpensesSubtotal(es) * TaxRate
}

// @todo: Refactor this so we're caching the results and sending them to the template rather than calling the functions in the template
func ExpensesWithTaxesSubtotal(es []Expense) float64 {
	return ExpensesSubtotal(es) + ExpensesTaxes(es)
}

func LabourSubtotal(ls []Labour) float64 {
	sum := 0.0
	for _, l := range ls {
		sum += l.TotalCost()
	}
	return sum
}

func ReceiptTotal(es []Expense, ls []Labour) float64 {
	return ExpensesWithTaxesSubtotal(es) + LabourSubtotal(ls)
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

type Receipt struct {
	Owner         Owner      `json:"owner"`
	BilledTo      BilledTo   `json:"billedTo"`
	ExpenseList   []Expense  `json:"expenseList"`
	LabourList    []Labour   `json:"labourList"`
	InvoiceNumber int        `json:"invoiceNumber"`
	InvoiceDate   tools.Date `json:"invoiceDate"`
}

func Create() {
	// Temp read from db.json file
	jsonFile, err := os.Open("./storage/db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var receipt Receipt
	err = json.Unmarshal(byteValue, &receipt)
	if err != nil {
		log.Fatal(err)
	}

	templateData := struct {
		ExpenseList               []Expense
		LabourList                []Labour
		BilledTo                  BilledTo
		Owner                     Owner
		InvoiceNumber             int
		InvoiceDate               tools.Date
		ExpensesSubtotal          func([]Expense) float64
		ExpensesTaxes             func([]Expense) float64
		ExpensesWithTaxesSubtotal func([]Expense) float64
		LabourSubtotal            func([]Labour) float64
		ReceiptTotal              func([]Expense, []Labour) float64
	}{
		ExpenseList:               receipt.ExpenseList,
		LabourList:                receipt.LabourList,
		BilledTo:                  receipt.BilledTo,
		Owner:                     receipt.Owner,
		InvoiceNumber:             receipt.InvoiceNumber,
		InvoiceDate:               receipt.InvoiceDate,
		ExpensesSubtotal:          ExpensesSubtotal,
		ExpensesTaxes:             ExpensesTaxes,
		ExpensesWithTaxesSubtotal: ExpensesWithTaxesSubtotal,
		LabourSubtotal:            LabourSubtotal,
		ReceiptTotal:              ReceiptTotal,
	}

	tools.CreatePdf("receipt.tmpl", "./receipt.pdf", templateData)
}
