package models

type Summary struct {
	Subtotal		float64		`json:"subtotal,omitempty"`
	ShippingCost	float64		`json:"shipping_cost,omitempty"`
	TotalTax		float64		`json:"total_tax,omitempty"`
	TotalCost		float64		`json:"total_cost,omitempty"`
}