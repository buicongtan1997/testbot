package models

type Element struct {
	Title         string        `json:"title,omitempty"`
	Subtitle      string        `json:"subtitle,omitempty"`
	ImageURL      string        `json:"image_url,omitempty"`
	Buttons       []Button      `json:"buttons,omitempty"`
	Quantity	  int64			`json:"quantity,omitempty"`
	Price		  float64		`json:"price,omitempty"`
	Currency	  string		`json:"currency,omitempty"`
}
