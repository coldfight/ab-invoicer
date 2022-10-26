package main

import (
	dbService "github.com/coldfight/ab-invoicer/internal/services/db_service"
	invoiceService "github.com/coldfight/ab-invoicer/internal/services/invoice_service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbService.CreateInitialDatabase()
	//dbService.SeedDatabase()
	invoice := invoiceService.GetFullInvoiceRecord(1)
	invoiceService.NewDocumentFromInvoice(invoice)
}
