package invoice_service

import (
	"fmt"
	"github.com/coldfight/ab-invoicer/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @todo: Instead of failing log.Fatal we should be returning it and handling it nicely on the "front-end"

func GetFullInvoiceRecord(invoiceNumber models.InvoiceNumber) models.Invoice {
	db, err := gorm.Open(sqlite.Open("storage/test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}

	var invoice models.Invoice
	db.
		Preload("Owner").
		Preload("Customer").
		Preload("Expenses").
		Preload("Labours").
		First(&invoice, "invoice_number = ?", invoiceNumber)

	return invoice
}

func GetInvoices() []models.Invoice {
	db, err := gorm.Open(sqlite.Open("storage/test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %s", err.Error()))
	}

	var invoices []models.Invoice
	result := db.
		Preload("Owner").
		Preload("Customer").
		Preload("Expenses").
		Preload("Labours").
		Find(&invoices)
	if result.Error != nil {
		return []models.Invoice{}
	}

	return invoices
}
