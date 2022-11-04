package invoice_repository

import (
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/services/db_service"
)

func GetById(id int) (models.Invoice, error) {
	var invoice models.Invoice

	db, err := db_service.GetConnection()
	if err != nil {
		return invoice, err
	}
	defer db.Close()

	// Grab invoice, owner, customer
	stmt, err := db.Prepare(`
SELECT i.invoiceNumber, i.invoiceDate,
       o.name, o.street, o.city, o.province, o.postalCode, o.phone, o.email,
       c.name, c.street, c.city, c.province, c.postalCode
FROM invoices i
LEFT JOIN owners o ON o.id = owner
LEFT JOIN customers c ON c.id = billedTo
WHERE i.id = ?
`)
	if err != nil {
		return invoice, err
	}
	defer stmt.Close()

	var invoiceDateString string
	err = stmt.QueryRow(id).Scan(
		&invoice.InvoiceNumber, &invoiceDateString,
		&invoice.Owner.Name, &invoice.Owner.Street, &invoice.Owner.City, &invoice.Owner.Province, &invoice.Owner.PostalCode, &invoice.Owner.Phone, &invoice.Owner.Email,
		&invoice.Customer.Name, &invoice.Customer.Street, &invoice.Customer.City, &invoice.Customer.Province, &invoice.Customer.PostalCode,
	)
	if err != nil {
		return invoice, err
	}
	invoice.InvoiceDate.SetFromString("2006-01-02", invoiceDateString)

	return invoice, err
}

func GetAll() ([]models.Invoice, error) {
	db, err := db_service.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Grab invoice, owner, customer
	rows, err := db.Query(`
SELECT i.invoiceNumber, i.invoiceDate,
       o.name, o.street, o.city, o.province, o.postalCode, o.phone, o.email,
       c.name, c.street, c.city, c.province, c.postalCode
FROM invoices i
LEFT JOIN owners o ON o.id = owner
LEFT JOIN customers c ON c.id = billedTo
ORDER BY date(i.invoiceDate) ASC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []models.Invoice
	for rows.Next() {
		var invoice models.Invoice
		var invoiceDateString string

		err = rows.Scan(&invoice.InvoiceNumber, &invoiceDateString,
			&invoice.Owner.Name, &invoice.Owner.Street, &invoice.Owner.City, &invoice.Owner.Province, &invoice.Owner.PostalCode, &invoice.Owner.Phone, &invoice.Owner.Email,
			&invoice.Customer.Name, &invoice.Customer.Street, &invoice.Customer.City, &invoice.Customer.Province, &invoice.Customer.PostalCode)
		if err != nil {
			return nil, err
		}
		invoice.InvoiceDate.SetFromString("2006-01-02", invoiceDateString)
		invoices = append(invoices, invoice)
	}

	return invoices, err
}
