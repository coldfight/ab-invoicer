package models

type Expense struct {
	Quantity    int
	Description string
	UnitPrice   float64
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
