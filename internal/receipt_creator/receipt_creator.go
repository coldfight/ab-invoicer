package receipt_creator

import (
	"github.com/coldfight/ab-invoicer/internal/tools"
	"time"
)

type Expense struct {
	Quantity    int
	Description string
	UnitPrice   float64
}

func (e Expense) TotalCost() float64 {
	return float64(e.Quantity) * e.UnitPrice
}

type Labour struct {
	Date        time.Time
	Description string
	Amount      float64
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

func LabourSubtotal(ls []Labour) float64 {
	sum := 0.0
	for _, l := range ls {
		sum += l.TotalCost()
	}
	return sum
}

type BilledTo struct {
	Name       string
	Street     string
	City       string
	Province   string
	PostalCode string
	Phone      string
}

type Owner struct {
	Name       string
	Street     string
	City       string
	Province   string
	PostalCode string
	Phone      string
	Email      string
}

func Create() {
	expenses := []Expense{
		{
			Quantity:    3,
			Description: "Lysol Aerosol",
			UnitPrice:   8,
		},
		{
			Quantity:    1,
			Description: "Windex",
			UnitPrice:   8,
		},
		{
			Quantity:    2,
			Description: "Toilet Bowl Cleaner",
			UnitPrice:   3.50,
		},
	}

	labours := []Labour{
		{
			Date:        time.Now(),
			Description: "Sanitize, clean, dust, vacuum, take out garbage, window cleaning",
			Amount:      100,
		},
		{
			Date:        time.Now(),
			Description: "Sanitize, clean, dust, vacuum, take out garbage, window cleaning",
			Amount:      100,
		},
		{
			Date:        time.Now(),
			Description: "Sanitize, clean, dust, vacuum, take out garbage, window cleaning",
			Amount:      100,
		},
		{
			Date:        time.Now(),
			Description: "Sanitize, clean, dust, vacuum, take out garbage, window cleaning",
			Amount:      100,
		},
	}

	billedTo := BilledTo{
		Name:       "Applewood Gospel Hall",
		Street:     "15 Dummy Street Area",
		City:       "Mississauga",
		Province:   "Ontario",
		PostalCode: "M4T 2T1",
		Phone:      "416.555.5555",
	}

	owner := Owner{
		Name:       "My name",
		Street:     "555 Mississauga Valley Blvd.",
		City:       "Mississauga",
		Province:   "Ontario",
		PostalCode: "M4T 2T1",
		Phone:      "416.555.5456",
		Email:      "myemail@example.com",
	}

	templateData := struct {
		Expenses         []Expense
		Labours          []Labour
		BilledTo         BilledTo
		Owner            Owner
		GetAbsPath       func(string) string
		AsCurrency       func(float64) string
		ExpensesSubtotal func([]Expense) float64
		LabourSubtotal   func([]Labour) float64
	}{
		Expenses:         expenses,
		Labours:          labours,
		BilledTo:         billedTo,
		Owner:            owner,
		GetAbsPath:       tools.FullFilePath,
		AsCurrency:       tools.Currency,
		ExpensesSubtotal: ExpensesSubtotal,
		LabourSubtotal:   LabourSubtotal,
	}

	tools.CreatePdf("./templates/receipt.tmpl", "./storage/receipt.pdf", templateData)
}
