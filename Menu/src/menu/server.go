package main

import (
	"fmt"
	"log"
	// "math"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	// "github.com/gorilla/handlers"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/rs/cors"
)

var database_server = "127.0.0.1:27017"
var database 			= "menu"
var collection = "Menu"


// MenuServer configures and returns a MenuServer instance.
func MenuServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	corsObj := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.Use(corsObj)
	n.UseHandler(mx)
	return n
}

// Menu Service API Routes
func initRoutes(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	router.HandleFunc("/menu/item", createItemHandler(formatter)).Methods("POST")
	router.HandleFunc("/menu/item/{itemId}", getItem(formatter)).Methods("GET")
	router.HandleFunc("/menu/items/{itemType}", getItemList(formatter)).Methods("GET")
	router.HandleFunc("/menu/item/{itemId}", updateItemHandler(formatter)).Methods("PUT")
	router.HandleFunc("/menu/item/{itemId}", deleteItemHandler(formatter)).Methods("DELETE")
}

// Error Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// Menu Service Health Check API

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API for creating an item in the menu
func createItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var menuItem MenuItem
		_ = json.NewDecoder(request.Body).Decode(&menuItem)
    	fmt.Println("Item Payload on create Item handler", menuItem)
    	uuid,_	 := uuid.NewV4()
    	menuItem.ItemId = uuid.String()
    	session, err := mgo.Dial(database_server)
        if err != nil {
						fmt.Println( "Error while connecting to mongo: ", err )
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
       defer session.Close()

       mongo_collection := session.DB(database).C(collection)

       var item Item;
       item.ItemId = menuItem.ItemId
			 item.ItemType = menuItem.ItemType
       item.ItemName = menuItem.ItemName
			 item.ItemSummary = menuItem.ItemSummary
			 item.ItemDescription = menuItem.ItemDescription
			 item.ItemAmount = menuItem.ItemAmount
			 item.ItemCalorieContent = menuItem.ItemCalorieContent
			 item.ItemAvailable = true
       error := mongo_collection.Insert(item)
       fmt.Println("Error while inserting a document: ", error)
       if error != nil {
            formatter.JSON(response, http.StatusInternalServerError, "Error while inserting a document")
            return
       }
		formatter.JSON(response, http.StatusOK, item)
	}
}

// API for getting list of all items in the menu

func getItemList(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		fmt.Println( "itemType params: ", params )
		var itemType string = params["itemType"]
		fmt.Println( "Item TYPE: ", itemType )

		session, err := mgo.Dial(database_server)
        if err != nil {
						fmt.Println( "Error while connecting to mongo: ", err )
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
        mongo_collection := session.DB(database).C(collection)
        var result []bson.M
        err = mongo_collection.Find(bson.M{"itemavailable" : true , "itemtype" : itemType}).All(&result)
        if err != nil {
            formatter.JSON(response, http.StatusNotFound, "No item found!!!")
            return
        }
		formatter.JSON(response, http.StatusOK, result)
	}
}

// API for getting an item from the menu

func getItem(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		fmt.Println( "Item ID params: ", params )
		var itemId string = params["itemId"]
		fmt.Println( "Item ID: ", itemId )
		session, err := mgo.Dial(database_server)
        if err != nil {
						fmt.Println( "Error while connecting to mongodb : ", err )
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
        mongo_collection := session.DB(database).C(collection)
        var result bson.M
        err = mongo_collection.Find(bson.M{"itemid" : itemId,"itemavailable": true}).One(&result)
				fmt.Println( "Item ID error: ", err )
        if err != nil {
            formatter.JSON(response, http.StatusNotFound, "No item found with given id!")
            return
        }
		formatter.JSON(response, http.StatusOK, result)
	}
}

// API for updating an item from the menu
func updateItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if (*request).Method == "OPTIONS" {
			fmt.Println("PREFLIGHT Request")
			return
		}
		var menuItem MenuItem
		_ = json.NewDecoder(request.Body).Decode(&menuItem)

    	fmt.Println("Item Payload ", menuItem)
			// params := mux.Vars(request)
			var itemId string = menuItem.ItemId
			fmt.Println( "Item ID params: ", menuItem.ItemId )
    	session, err := mgo.Dial(database_server)
        if err != nil {
						fmt.Println( "Error while connecting to mongo: ", err )
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
       mongo_collection := session.DB(database).C(collection)

       	var item Item;
        err = mongo_collection.Find(bson.M{"itemid" : itemId}).One(&item)
				if err != nil {
									fmt.Println("error: ", err)
									formatter.JSON(response, http.StatusNotFound, "Item not found")
								return
							}else{
							if item.ItemId == itemId {
								item.ItemName = menuItem.ItemName
								item.ItemType = menuItem.ItemType
								item.ItemSummary = menuItem.ItemSummary
								item.ItemDescription = menuItem.ItemDescription
								item.ItemAmount = menuItem.ItemAmount
								item.ItemCalorieContent = menuItem.ItemCalorieContent
								item.ItemAvailable = true
							}
								error := mongo_collection.Update(bson.M{"itemid": item.ItemId}, bson.M{"$set": bson.M{"itemname": item.ItemName,
									"itemtype": item.ItemType,"itemsummary": item.ItemSummary,"itemdescription": item.ItemDescription,"itemamount": item.ItemAmount,
									"itemcalorieContent": item.ItemCalorieContent,"itemavailable" : true }})

								if error != nil {
									fmt.Println("error: ", error)
											formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
											return
								}
							}

					formatter.JSON(response, http.StatusOK, item)
				}
}

// API to delete an item from menu
func deleteItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var menuItem DeleteMenuItem
		_ = json.NewDecoder(request.Body).Decode(&menuItem)
    	fmt.Println("Item Payload ", menuItem)

			params := mux.Vars(request)
			fmt.Println( "Item ID params: ", params )
			var itemId string = params["itemId"]
			fmt.Println( "Item ID params: ", itemId )
    	session, err := mgo.Dial(database_server)
        if err != nil {
						fmt.Println( "Error while connecting to mongodb: ", err )
            formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
            return
        }
        defer session.Close()
        mongo_collection := session.DB(database).C(collection)

       	var item Item;
        err = mongo_collection.Find(bson.M{"itemid" : itemId}).One(&item)
        if err != nil {
            fmt.Println("error: ", err)
            formatter.JSON(response, http.StatusNotFound, "No item found with given id !!!")
        	return
        }else{
        	error := mongo_collection.Update(bson.M{"itemid": itemId}, bson.M{"$set": bson.M{"itemavailable": false}})
        	if error != nil {
        		fmt.Println("error: ", error)
                formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
                return
        	}
        }
		formatter.JSON(response, http.StatusOK, item)
	}
}
