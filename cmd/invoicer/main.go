package main

import (
	"github.com/coldfight/ab-invoicer/internal/ui/list_invoices"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//dbService.CreateInitialDatabase()
	//invoice := invoiceService.GetFullInvoiceRecord(1)
	//invoiceService.NewDocumentFromInvoice(invoice)

	list_invoices.Run()
}
