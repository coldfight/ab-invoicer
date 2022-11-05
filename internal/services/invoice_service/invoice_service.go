package invoice_service

import (
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/services/db_service"
)

// @todo: Instead of failing log.Fatal we should be returning it and handling it nicely on the "front-end"

func GetFullInvoiceRecord(invoiceNumber models.InvoiceNumber) models.Invoice {
	db := db_service.GetConnection()

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
	db := db_service.GetConnection()

	var invoices []models.Invoice
	result := db.
		Preload("Owner").
		Preload("Customer").
		Preload("Expenses").
		Preload("Labours").
		Order("invoice_number asc").
		Find(&invoices)

	if result.Error != nil {
		return []models.Invoice{}
	}

	return invoices
}
