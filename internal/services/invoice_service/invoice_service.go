package invoice_service

import (
	"database/sql"
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/repositories/expense_repository"
	"github.com/coldfight/ab-invoicer/internal/repositories/labour_repository"
	"github.com/coldfight/ab-invoicer/internal/services/db_service"
	"log"
)

func GetFullInvoiceRecord(id int) models.Invoice {
	db, err := sql.Open(db_service.DriverName, db_service.DbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var invoice models.Invoice

	// Grab invoice, owner, customer
	stmt, err := db.Prepare(`
SELECT i.id, i.invoiceDate,
       o.name, o.street, o.city, o.province, o.postalCode, o.phone, o.email,
       c.name, c.street, c.city, c.province, c.postalCode
FROM invoices i
LEFT JOIN owners o ON o.id = owner
LEFT JOIN customers c ON c.id = billedTo
WHERE i.id = ?
`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var invoiceDateString string
	err = stmt.QueryRow(id).Scan(
		&invoice.InvoiceNumber, &invoiceDateString,
		&invoice.Owner.Name, &invoice.Owner.Street, &invoice.Owner.City, &invoice.Owner.Province, &invoice.Owner.PostalCode, &invoice.Owner.Phone, &invoice.Owner.Email,
		&invoice.BilledTo.Name, &invoice.BilledTo.Street, &invoice.BilledTo.City, &invoice.BilledTo.Province, &invoice.BilledTo.PostalCode,
	)
	if err != nil {
		log.Fatal(err)
	}
	invoice.InvoiceDate.SetFromString("2006-01-02", invoiceDateString)

	// Grab expenses
	expenseList, err := expense_repository.GetByInvoiceId(id, db)
	if err != nil {
		log.Fatal(err)
	}
	invoice.ExpenseList = expenseList

	// Grab labour
	labourList, err := labour_repository.GetByInvoiceId(id, db)
	if err != nil {
		log.Fatal(err)
	}
	invoice.LabourList = labourList

	return invoice
}
