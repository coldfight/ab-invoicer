package models

import (
	"fmt"
	"time"
)

const (
	TaxRate = 0.13
)

type Date time.Time
type InvoiceNumber int

func (num InvoiceNumber) Padded() string {
	return fmt.Sprintf("%03d", num)
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

type Invoice struct {
	Owner         Owner
	BilledTo      Customer
	ExpenseList   ExpenseList
	LabourList    LabourList
	InvoiceNumber InvoiceNumber
	InvoiceDate   Date
}

func (i Invoice) Total() float64 {
	return i.ExpenseList.ExpensesWithTaxesSubtotal() + i.LabourList.LabourSubtotal()
}
