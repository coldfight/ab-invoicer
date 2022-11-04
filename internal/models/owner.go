package models

import "gorm.io/gorm"

type Owner struct {
	gorm.Model
	Invoices   []Invoice // Has many
	Name       string
	Street     string
	City       string
	Province   string
	PostalCode string
	Phone      string
	Email      string
}
