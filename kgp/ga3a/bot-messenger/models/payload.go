package models

type Payload struct {
	URL               	string    		`json:"url,omitempty"`
	TemplateType      	string    		`json:"template_type,omitempty"`
	Sharable         	bool      		`json:"sharable,omitempty"`
	ImageAspectRation 	string    		`json:"image_aspect_ratio,omitempty"`
	Elements          	[]Element 		`json:"elements,omitempty"`
	RecipientName    	string			`json:"recipient_name,omitempty"`
	OrderNumber 		string 			`json:"order_number,omitempty"`
	Currency			string			`json:"currency,omitempty"`
	PaymentMethod		string 			`json:"payment_method,omitempty"`
	OrderUrl			string 			`json:"order_url,omitempty"`
	Timestamp			int64			`json:"timestamp,omitempty"`
	Address				Address			`json:"address,omitempty"`
	Summary				Summary			`json:"summary,omitempty"`
	Adjustments			[]Adjustment	`json:"adjustments,omitempty"`

}
