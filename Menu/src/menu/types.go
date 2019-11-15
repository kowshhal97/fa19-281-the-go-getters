package main

// Menu Item structure
type Item struct {
	ItemId							string
	ItemType						string
	ItemName		    		    string
	ItemSummary 				    string
	ItemDescription 		        string
	ItemAmount		    	        float64	`json:",string"`
	ItemCalorieContent []ItemCalorie
	ItemAvailable				     bool
}

type MenuItem struct {

	ItemId							 string
	ItemType						 string
	ItemName		    		     string
	ItemSummary 				     string
	ItemDescription 		         string
	ItemAmount		    	         float64	`json:",string"`
	ItemCalorieContent	[]ItemCalorie
	ItemImagePath   		         string
	ItemAvailable				     bool
}

type DeleteMenuItem struct {
	ItemId							string
}

type ItemCalorie struct {
	Content            string
	Amount             string
}
