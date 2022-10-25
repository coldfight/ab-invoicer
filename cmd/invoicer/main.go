package main

import (
	"encoding/json"
	"github.com/coldfight/ab-invoicer/internal/invoice_generator"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"os"
)

func main() {
	//invoice_db_service.CreateInvoiceDatabase()
	//invoice_db_service.SeedDatabase()
	//invoice_db_service.GetFullInvoiceRecord(1)

	// Temp read from db.json file
	jsonFile, err := os.Open("./storage/db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var invoice invoice_generator.Invoice
	err = json.Unmarshal(byteValue, &invoice)
	if err != nil {
		log.Fatal(err)
	}
	invoice_generator.NewInvoice(invoice)
}
