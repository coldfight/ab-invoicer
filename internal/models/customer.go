package models

type Customer struct {
	Name       string `json:"name"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postalCode"`
	Phone      string `json:"phone"`
}
