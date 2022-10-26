package main

import (
	"github.com/coldfight/ab-invoicer/internal/invoice_generator"
	"github.com/coldfight/ab-invoicer/internal/services/invoice_service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//db_service.CreateInitialDatabase()
	//db_service.SeedDatabase()
	invoice := invoice_service.GetFullInvoiceRecord(1)

	invoice_generator.NewInvoice(invoice)
}
