package models

import "gorm.io/gorm"

type Owner struct {
	Name       string
	Street     string
	City       string
	Province   string
	PostalCode string
	Phone      string
	Email      string
	// --
	gorm.Model
	Invoices []Invoice // Has many
}
