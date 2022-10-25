package db

import (
	"database/sql"
	"log"
)

const (
	DriverName = "sqlite3"
	DbFilePath = "./storage/invoices.db"
)

// @todo: Hardcoded for invoices right now....
func CreateDatabase() {
	// @todo: This is temporary... we don't want to delete this every time we load up the application
	//os.Remove(DbFilePath)

	db, err := sql.Open(DriverName, DbFilePath)
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

	//_, err = db.Exec(sqlStmt)
	//if err != nil {
	//	log.Printf("%q: %s\n", err, sqlStmt)
	//	return
	//}
	//
	//tx, err := db.Begin()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//stmt, err := tx.Prepare("INSERT INTO foo(id, name) values(?, ?)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()
	//
	//for i := 0; i < 100; i++ {
	//	_, err = stmt.Exec(i, fmt.Sprintf("Hello World; こんにちは世界; %03d", i))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//
	//err = tx.Commit()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//rows, err := db.Query("SELECT id, name FROM foo")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	var id int
	//	var name string
	//	err = rows.Scan(&id, &name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(id, name)
	//}
	//
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//stmt, err = db.Prepare("SELECT name from foo where id = ?")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()
	//var name string
	//err = stmt.QueryRow("3").Scan(&name)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(name)
	//
	//_, err = db.Exec("DELETE FROM foo")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = db.Exec("INSERT INTO foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//rows, err = db.Query("SELECT id, name FROM foo")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	var id int
	//	var name string
	//	err = rows.Scan(&id, &name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(id, name)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
