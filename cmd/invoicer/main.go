package main

import (
	"encoding/json"
	"github.com/coldfight/ab-invoicer/internal/invoice_generator"
	"io"
	"log"
	"os"
)

func main() {
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
