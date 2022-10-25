package invoice_db_service

import (
	"database/sql"
	"fmt"
	"github.com/coldfight/ab-invoicer/internal/invoice_generator"
	"log"
)

const (
	driverName = "sqlite3"
	dbFilePath = "./storage/invoices.db"
)

func GetFullInvoiceRecord(id int) {
	db, err := sql.Open(driverName, dbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
SELECT i.id, -- i.invoiceDate,
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

	err = stmt.QueryRow(id).Scan(
		&invoice.InvoiceNumber, // &invoice.InvoiceDate,
		&invoice.Owner.Name, &invoice.Owner.Street, &invoice.Owner.City, &invoice.Owner.Province, &invoice.Owner.PostalCode, &invoice.Owner.Phone, &invoice.Owner.Email,
		&invoice.BilledTo.Name, &invoice.BilledTo.Street, &invoice.BilledTo.City, &invoice.BilledTo.Province, &invoice.BilledTo.PostalCode,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(invoice)

}

func CreateInvoiceDatabase() {
	// @todo: This is temporary... we don't want to delete this every time we load up the application
	//os.Remove(dbFilePath)

	db, err := sql.Open(driverName, dbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createOwnersStmt := `
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

	createCustomersStmt := `
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

	createInvoicesStmt := `
CREATE TABLE IF NOT EXISTS invoices
(
    id          INTEGER NOT NULL PRIMARY KEY,
    owner       INTEGER NOT NULL DEFAULT 0,
    billedTo    INTEGER NOT NULL DEFAULT 0,
    invoiceDate TEXT    NOT NULL DEFAULT '',
    FOREIGN KEY (owner) REFERENCES owners (id) ON DELETE CASCADE,
    FOREIGN KEY (billedTo) REFERENCES customers (id) ON DELETE CASCADE
);`

	createExpensesStmt := `
CREATE TABLE IF NOT EXISTS expenses
(
    id          INTEGER NOT NULL PRIMARY KEY,
    invoice     INTEGER NOT NULL,
    quantity    INTEGER,
    description TEXT,
    unitPrice   REAL,
    FOREIGN KEY (invoice) REFERENCES invoices (id) ON DELETE CASCADE
);`

	createLaboursStmt := `
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

	tx.Exec(createOwnersStmt)
	tx.Exec(createCustomersStmt)
	tx.Exec(createInvoicesStmt)
	tx.Exec(createExpensesStmt)
	tx.Exec(createLaboursStmt)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
