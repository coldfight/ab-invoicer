package db_service

import (
	"database/sql"
	"log"
)

const (
	DriverName = "sqlite3"
	DbFilePath = "./storage/invoices.db"
)

func GetConnection() (*sql.DB, error) {
	return sql.Open(DriverName, DbFilePath)
}

func CreateInitialDatabase() {
	// @todo: This is temporary... we don't want to delete this every time we load up the application
	//os.Remove(DbFilePath)

	db, err := GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
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
)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
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
)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
CREATE TABLE IF NOT EXISTS invoices
(
    id          INTEGER NOT NULL PRIMARY KEY,
    owner       INTEGER NOT NULL DEFAULT 0,
    billedTo    INTEGER NOT NULL DEFAULT 0,
    invoiceDate TEXT    NOT NULL DEFAULT '',
    FOREIGN KEY (owner) REFERENCES owners (id) ON DELETE CASCADE,
    FOREIGN KEY (billedTo) REFERENCES customers (id) ON DELETE CASCADE
)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
CREATE TABLE IF NOT EXISTS expenses
(
    id          INTEGER NOT NULL PRIMARY KEY,
    invoice     INTEGER NOT NULL,
    quantity    INTEGER,
    description TEXT,
    unitPrice   REAL,
    FOREIGN KEY (invoice) REFERENCES invoices (id) ON DELETE CASCADE
)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
CREATE TABLE IF NOT EXISTS labour
(
    id          INTEGER NOT NULL PRIMARY KEY,
    invoice     INTEGER NOT NULL,
    date        TEXT,
    description TEXT,
    amount      REAL,
    FOREIGN KEY (invoice) REFERENCES invoices (id) ON DELETE CASCADE
)`)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
