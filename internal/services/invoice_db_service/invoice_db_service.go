package invoice_db_service

import (
	"database/sql"
	"github.com/coldfight/ab-invoicer/internal/invoice_generator"
	"log"
)

const (
	DriverName = "sqlite3"
	DbFilePath = "./storage/invoices.db"
)

func GetFullInvoiceRecord(id int) invoice_generator.Invoice {
	db, err := sql.Open(DriverName, DbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
SELECT i.id, i.invoiceDate,
       o.name, o.street, o.city, o.province, o.postalCode, o.phone, o.email,
       c.name, c.street, c.city, c.province, c.postalCode
FROM invoices i
LEFT JOIN owners o ON o.id = owner
LEFT JOIN customers c ON c.id = billedTo
WHERE i.id = ?
`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var invoice invoice_generator.Invoice

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

	return invoice
}

func CreateInvoiceDatabase() {
	// @todo: This is temporary... we don't want to delete this every time we load up the application
	//os.Remove(DbFilePath)

	db, err := sql.Open(DriverName, DbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createOwnersQuery := `
CREATE TABLE IF NOT EXISTS owners
(
    id         INTEGER NOT NULL PRIMARY KEY,
    name       TEXT    NOT NULL DEFAULT '',
    street     TEXT    NOT NULL DEFAULT '',
    city       TEXT    NOT NULL DEFAULT '',
    province   TEXT    NOT NULL DEFAULT '',
    postalCode TEXT    NOT NULL DEFAULT '',
    phone      TEXT    NOT NULL DEFAULT '',
    email      TEXT    NOT NULL DEFAULT ''
);`

	createCustomersQuery := `
CREATE TABLE IF NOT EXISTS customers
(
    id         INTEGER NOT NULL PRIMARY KEY,
    name       TEXT    NOT NULL DEFAULT '',
    street     TEXT    NOT NULL DEFAULT '',
    city       TEXT    NOT NULL DEFAULT '',
    province   TEXT    NOT NULL DEFAULT '',
    postalCode TEXT    NOT NULL DEFAULT '',
    phone      TEXT    NOT NULL DEFAULT '',
    email      TEXT    NOT NULL DEFAULT ''
);`

	createInvoicesQuery := `
CREATE TABLE IF NOT EXISTS invoices
(
    id          INTEGER NOT NULL PRIMARY KEY,
    owner       INTEGER NOT NULL DEFAULT 0,
    billedTo    INTEGER NOT NULL DEFAULT 0,
    invoiceDate TEXT    NOT NULL DEFAULT '',
    FOREIGN KEY (owner) REFERENCES owners (id) ON DELETE CASCADE,
    FOREIGN KEY (billedTo) REFERENCES customers (id) ON DELETE CASCADE
);`

	createExpensesQuery := `
CREATE TABLE IF NOT EXISTS expenses
(
    id          INTEGER NOT NULL PRIMARY KEY,
    invoice     INTEGER NOT NULL,
    quantity    INTEGER,
    description TEXT,
    unitPrice   REAL,
    FOREIGN KEY (invoice) REFERENCES invoices (id) ON DELETE CASCADE
);`

	createLaboursQuery := `
CREATE TABLE IF NOT EXISTS labour
(
    id          INTEGER NOT NULL PRIMARY KEY,
    invoice     INTEGER NOT NULL,
    date        TEXT,
    description TEXT,
    amount      REAL,
    FOREIGN KEY (invoice) REFERENCES invoices (id) ON DELETE CASCADE
);`

	tx, err := db.Begin()

	tx.Exec(createOwnersQuery)
	tx.Exec(createCustomersQuery)
	tx.Exec(createInvoicesQuery)
	tx.Exec(createExpensesQuery)
	tx.Exec(createLaboursQuery)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
