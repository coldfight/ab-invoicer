package models

import "gorm.io/gorm"

type Customer struct {
	Name       string
	Street     string
	City       string
	Province   string
	PostalCode string
	Phone      string
	// --
	gorm.Model
	Invoices []Invoice // Has many
}
