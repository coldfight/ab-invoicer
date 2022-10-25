package main

import (
	"github.com/coldfight/ab-invoicer/internal/invoice_generator"
	"github.com/coldfight/ab-invoicer/internal/services/invoice_db_service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//invoice_db_service.CreateInvoiceDatabase()
	//invoice_db_service.SeedDatabase()
	invoice := invoice_db_service.GetFullInvoiceRecord(1)

	invoice_generator.NewInvoice(invoice)
}
