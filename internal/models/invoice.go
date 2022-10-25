package models

import (
	"encoding/json"
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

type Invoice struct {
	Owner         Owner       `json:"owner"`
	BilledTo      Customer    `json:"billedTo"`
	ExpenseList   ExpenseList `json:"expenseList"`
	LabourList    LabourList  `json:"labourList"`
	InvoiceNumber int         `json:"invoiceNumber"`
	InvoiceDate   Date        `json:"invoiceDate"`
}

func (i Invoice) Total() float64 {
	return i.ExpenseList.ExpensesWithTaxesSubtotal() + i.LabourList.LabourSubtotal()
}
