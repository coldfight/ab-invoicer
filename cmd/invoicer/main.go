package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/ui/app"
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

	//db_service.CreateInitialDatabase()
	//invoice := invoice_service.GetFullInvoiceRecord(3)
	//invoice_service.NewDocumentFromInvoice(invoice)

	app.Run()
}
