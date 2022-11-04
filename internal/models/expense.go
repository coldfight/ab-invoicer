package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	InvoiceID   uint // foreign key
	Quantity    int
	Description string
	UnitPrice   float64 // @todo: Update this to custom type `type Currency float64` so we can have a formatted() method
}
type ExpenseList []Expense

func (e Expense) TotalCost() float64 {
	return float64(e.Quantity) * e.UnitPrice
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

func (el ExpenseList) ExpensesWithTaxesSubtotal() float64 {
	return el.ExpensesSubtotal() + el.ExpensesTaxes()
}
