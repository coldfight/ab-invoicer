package receipt_creator

import (
	"encoding/json"
	"github.com/coldfight/ab-invoicer/internal/tools"
	"io"
	"log"
	"os"
	"time"
)

const (
	SavedDateLayout = "Jan 02, 2006"
)

type Expense struct {
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unitPrice"`
}

func (e Expense) TotalCost() float64 {
	return float64(e.Quantity) * e.UnitPrice
}

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

func (d *Date) MarshalJSON() ([]byte, error) {
	// @todo: I might want to call the json.Marshal instead of manually appending the `"`
	return []byte(`"` + time.Time(*d).Format(SavedDateLayout) + `"`), nil
}

func (d *Date) Format(layout string) string {
	return time.Time(*d).Format(layout)
}

type Labour struct {
	Date        Date    `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
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
	Owner       Owner     `json:"owner"`
	BilledTo    BilledTo  `json:"billedTo"`
	ExpenseList []Expense `json:"expenseList"`
	LabourList  []Labour  `json:"labourList"`
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
		ExpenseList      []Expense
		LabourList       []Labour
		BilledTo         BilledTo
		Owner            Owner
		GetAbsPath       func(string) string
		AsCurrency       func(float64) string
		ExpensesSubtotal func([]Expense) float64
		LabourSubtotal   func([]Labour) float64
	}{
		ExpenseList:      receipt.ExpenseList,
		LabourList:       receipt.LabourList,
		BilledTo:         receipt.BilledTo,
		Owner:            receipt.Owner,
		GetAbsPath:       tools.FullFilePath,
		AsCurrency:       tools.Currency,
		ExpensesSubtotal: ExpensesSubtotal,
		LabourSubtotal:   LabourSubtotal,
	}

	tools.CreatePdf("./templates/receipt.tmpl", "./storage/receipt.pdf", templateData)
}
