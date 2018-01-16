package models

type Address struct {
	Street1		string			`json:"street_1,omitempy"`
	Street2		string			`json:"street_2,omitempy"`
	City			string			`json:"city,omitempty"`
	PostalCode		string			`json:"postal_code,omitempty"`
	State			string			`json:"state,omitempty"`
	Country			string			`json:"country,omitempy"`
}