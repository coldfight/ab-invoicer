package db_service

import (
	"database/sql"
	"log"
)

func SeedDatabase() {
	db, err := sql.Open(DriverName, DbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
INSERT INTO customers(id, name, street, city, province, postalCode, phone) 
VALUES(1, 'Customer Namehere', '123 Duncan Rd', 'Mississauga', 'Ontario', 'M4T 2T1', '905.555.5555')
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
INSERT INTO owners(id, name, street, city, province, postalCode, phone, email) 
VALUES(1, 'Owner Namehere', '321 Kinsey Cres', 'Mississauga', 'Ontario', 'M4T 2T1', '416.555.5555', 'example@example.com')
`)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
INSERT INTO invoices(id, owner, billedTo, invoiceDate) 
VALUES(1, 1, 1, '2022-09-23')
`)
	if err != nil {
		log.Fatal(err)
	}

	expensesStmt, err := tx.Prepare(`
INSERT INTO expenses(id, invoice, quantity, description, unitPrice)
VALUES(?, ?, ?, ?, ?)
`)
	if err != nil {
		log.Fatal(err)
	}
	defer expensesStmt.Close()

	expenseRecords := []struct {
		id          int
		invoice     int
		quantity    int
		description string
		amount      float64
	}{
		{1, 1, 3, "Lysol Aerosol Spray", 8.99},
		{2, 1, 1, "Windex", 4.99},
		{3, 1, 1, "Paper Towel Rolls", 18.99},
	}

	for _, r := range expenseRecords {
		_, err = expensesStmt.Exec(r.id, r.invoice, r.quantity, r.description, r.amount)
		if err != nil {
			log.Fatal("something bad happened:", err)
		}
	}

	labourStmt, err := tx.Prepare(`
INSERT INTO labour(id, invoice, date, description, amount)
VALUES(?, ?, ?, ?, ?)
`)
	if err != nil {
		log.Fatal(err)
	}
	defer labourStmt.Close()

	labourRecords := []struct {
		id          int
		invoice     int
		date        string
		description string
		amount      float64
	}{
		{1, 1, "2022-09-03", "Sanitize, clean, dust, vacuum, take out garbage, window cleaning", 100},
		{2, 1, "2022-09-10", "Sanitize, clean, dust, vacuum, take out garbage, window cleaning", 100},
		{3, 1, "2022-09-17", "Sanitize, clean, dust, vacuum, take out garbage, window cleaning", 100},
		{4, 1, "2022-09-24", "Sanitize, clean, dust, vacuum, take out garbage, window cleaning", 100},
	}

	for _, r := range labourRecords {
		_, err = labourStmt.Exec(r.id, r.invoice, r.date, r.description, r.amount)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
