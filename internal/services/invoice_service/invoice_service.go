package invoice_service

import (
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/repositories/expense_repository"
	"github.com/coldfight/ab-invoicer/internal/repositories/invoice_repository"
	"github.com/coldfight/ab-invoicer/internal/repositories/labour_repository"
	"log"
)

func GetFullInvoiceRecord(id int) models.Invoice {
	// Grab main invoice
	invoice, err := invoice_repository.GetById(id)
	if err != nil {
		log.Fatal(err)
	}

	// Grab expenses
	expenseList, err := expense_repository.GetByInvoiceId(id)
	if err != nil {
		log.Fatal(err)
	}
	invoice.ExpenseList = expenseList

	// Grab labour
	labourList, err := labour_repository.GetByInvoiceId(id)
	if err != nil {
		log.Fatal(err)
	}
	invoice.LabourList = labourList

	return invoice
}

// @todo: Instead of failing log.Fatal we should be returning it and handling it nicely on the "front-end"

func GetInvoices() []models.Invoice {
	invoices, err := invoice_repository.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	return invoices
}
