/*
	Pizza ordering APIs
*/
	
package main


type PizzaOrder struct {
	OrderId     string       `json:"orderId" bson:"orderId"`
	UserId      string       `json:"userId" bson:"userId"`
	OrderStatus string       `json:"orderStatus" bson:"orderStatus"`
	Items        []PizzaItem `json:"items" bson:"items"`
	TotalAmount float32      `json:"totalAmount" bson:"totalAmount"`
}

type RequiredPayload struct {
	OrderId     string  `json:"orderId" bson:"orderId"`
	UserId      string  `json:"userId" bson:"userId"`
	ItemName    string  `json:"itemName" bson:"itemId"`
	ItemId      string  `json:"itemId" bson:"itemId"`
	ItemPrice       float32 `json:"itemPrice" bson:"itemPrice"`
	ItemQuantity    float32  `json:"itemQuantity" bson:"itemQuantity"`
}

type PizzaItem struct {
	ItemName    string  `json:"itemName" bson:"itemName"`
	ItemId      string  `json:"itemId" bson:"itemId"`
	ItemPrice   float32 `json:"itemPrice" bson:"itemPrice"`
	ItemQuantity    float32  `json:"itemQuantity" bson:"itemQuantity"`
}


var orders map[string]PizzaOrder

