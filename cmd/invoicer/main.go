package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/services/db_service"
	"github.com/coldfight/ab-invoicer/internal/ui/app"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	if f, err := tea.LogToFile("logs/debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	// @todo: convert to using Gorm
	db_service.CreateInitialDatabase()
	//invoice := invoiceService.GetFullInvoiceRecord(1)
	//invoiceService.NewDocumentFromInvoice(invoice)

	app.Run()
}
