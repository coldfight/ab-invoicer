package models

import "gorm.io/gorm"

type Owner struct {
	gorm.Model
	Name       string
	Street     string
	City       string
	Province   string
	PostalCode string
	Phone      string
	Email      string
}
