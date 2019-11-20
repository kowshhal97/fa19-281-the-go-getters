package main

type Order struct {
	item     string
	quantity int
	price    float32
}

type Payment struct {
	PaymentID   string  //`json:"paymentId", mitempty"`
	OrderID     string  `json:"orderId",omitempty"`
	UserID      string  `json:"userId",omitempty"`
	TotalPrice  float32 `json:"totalPrice",omitempty"`
	CardDetails string  `json:"cardDetails",omitempty"`
	Contactphone string  `json:"contactPhone",omitempty"`
	Expdate string  	`json:"expDate",omitempty"`
	SecurityCode string `json:"securityCode",omitempty"`
	OrderStatus bool    //`json:"status,omitempty"`
	PaymentDate string  //`json:"paymentDate,omitempty"`
}