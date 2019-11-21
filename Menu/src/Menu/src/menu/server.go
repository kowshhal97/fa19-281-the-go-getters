package main

import (
	"fmt"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net"
	//	"log"
	"net/http"
	"strings"
	//	"os"
)

var database_server = "35.166.110.191:27017"
var database = "admin"
var collection = "menu"
var mongo_user = "admin"
var mongo_pass = "admin"

// MenuServer configures and returns a MenuServer instance.
func MenuServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD","DELETE","OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders,allowedMethods , allowedOrigins)(router))
	return n
}

// Menu Service API Routes
func initRoutes(router *mux.Router, formatter *render.Render) {
	router.HandleFunc("/menu/ping", pingHandler(formatter)).Methods("GET")
	router.HandleFunc("/menu", GetMenu(formatter)).Methods("GET")
	router.HandleFunc("/menu/{ItemId}", getItemByIdHandler(formatter)).Methods("GET")
	router.HandleFunc("/menu/item", createItemHandler(formatter)).Methods("POST")
}

/* Error Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}*/

// Menu Service Health Check API
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Pizza My Heart API Server is UP!: " + getSystemIp()
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

func getSystemIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	address := strings.Split(localAddr, ":")
	fmt.Println("address: ", address[0])
	return address[0]
}

// API to find an item in the menu
func GetMenu(formatter *render.Render) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("Here in Get")
		session,err:= mgo.Dial(database_server)

		//defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		fmt.Println("user pass", mongo_user, mongo_pass)
		err= session.DB(database).Login(mongo_user, mongo_pass)
		defer session.Close()
		fmt.Println("Defer session close")
		//session.SetMode(mgo.Monotonic, true) need to check-original comment
		mongo_collection := session.DB(database).C(collection)
		var result []Item
		fmt.Println("Declaring err")
		err = mongo_collection.Find(bson.M{}).All(&result)
		fmt.Println("Result: ", result)
		if err != nil {
			//log.Fatalf(" %s", err)
			formatter.JSON(response, http.StatusNotFound, "Menu not found !!!")
			return
		}
		fmt.Println("Menu : ", result)
		formatter.JSON(response, http.StatusOK, result)
	}

}


func getItemByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//Setup
		session, error := mgo.Dial(database_server)
		if error = session.DB(database).Login(mongo_user, mongo_pass); error != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Error")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)

		c := session.DB(database).C(collection)
		fmt.Println("Session established")
		parameter := mux.Vars(req)
		var item Item
		var item_id string = parameter["ItemId"]

		item_err := c.Find(bson.M{"ItemId": item_id}).One(&item)
		if item_err != nil {
			formatter.JSON(w, http.StatusNotFound, " This item id doesn't exist.")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&item)
		fmt.Println("Item details: ", item)
		formatter.JSON(w, http.StatusOK, item)
	}
}

func createItemHandler(formatter *render.Render) http.HandlerFunc{
	return func(response http.ResponseWriter, req *http.Request){

		var create_item MenuItem
		_ = json.NewDecoder(req.Body).Decode(&create_item)
		session, error := mgo.Dial(database_server)
		if error = session.DB(database).Login(mongo_user, mongo_pass);
		error != nil {
			formatter.JSON(response, http.StatusInternalServerError, "Error")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(database).C(collection)

		var new_item Item

		new_item.ItemId = create_item.ItemId
		new_item.ItemName = create_item.ItemName
		new_item.Price = create_item.Price
		new_item.Description = create_item.Description
        new_item.ItemType = create_item.ItemType

		item_error := c.Insert(new_item)
			if item_error != nil {
				formatter.JSON(response, http.StatusInternalServerError, "Sorry, item cannot be inserted, check values.")
				return
			}
			fmt.Println("New item inserted successfully.")
			formatter.JSON(response, http.StatusOK, new_item)
    }
}