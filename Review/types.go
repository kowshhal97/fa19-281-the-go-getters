package main

type Review struct {
	ReviewId   string  //`json:"reviewId",omitempty"`
	UserId     string  `json:"userId",omitempty" bson:"userId"`
	MenuItemId string  `json:"menuItem",omitempty" bson:"menuItem"`
	ReviewDate string  //`json:"paymentDate,omitempty"`
	Comment    string  `json:"comment,omitempty" bson:"comment"`
}
//var orders map[string] order

