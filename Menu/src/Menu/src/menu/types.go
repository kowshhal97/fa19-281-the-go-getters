package main

//Item struct
type Item struct {
	ItemId          string 	`json:"ItemId" bson:"ItemId"`
	ItemName		string  `json:"ItemName" bson:"ItemName"`
	Price 			int	    `json:"Price" bson:"Price"`
	Description 	string	`json:"Description" bson:"Description"`
	ItemType		string  `json:"ItemType" bson:"ItemType"`
}