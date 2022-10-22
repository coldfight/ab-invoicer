package receipt_creator

import (
	"github.com/coldfight/ab-invoicer/internal/tools"
	"time"
)

//type Totalable interface {
//	TotalCost() float64
//}

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

//func Subtotal(t []Totalable) float64 {
//	for i, e := range t {
//		fmt.Println(i, e)
//	}
//	return 0.0
//}

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

	templateData := struct {
		Expenses         []Expense
		Labours          []Labour
		GetAbsPath       func(string) string
		ExpensesSubtotal func([]Expense) float64
		LabourSubtotal   func([]Labour) float64
	}{
		Expenses:         expenses,
		Labours:          labours,
		GetAbsPath:       tools.FullFilePath,
		ExpensesSubtotal: ExpensesSubtotal,
		LabourSubtotal:   LabourSubtotal,
	}

	tools.CreatePdf("./templates/receipt.tmpl", "./storage/receipt.pdf", templateData)
}
