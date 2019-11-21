package main

import (
	"fmt"
//	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/gorilla/handlers"
	"gopkg.in/mgo.v2"
	"net"
	"strings"
	"gopkg.in/mgo.v2/bson"
//	"os"
)

// MongoDB Config
// var database_server = "ds227185.mlab.com:27185"
// var database = "counterburger"
// var collection = "menu"
// var mongo_user = "cmpe281"
// var mongo_pass = "cmpe281"

/*
var database_server = os.Getenv("Server")
var database = os.Getenv("Database")
var collection = os.Getenv("Collection")
var mongo_user = os.Getenv("User")
var mongo_pass = os.Getenv("Pass")
*/

var database_server = "35.166.110.191:27017"
var database = "admin"
var collection = "users"
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
		if(err!=nil){
		formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error*****")
        			return
		}
		//defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		fmt.Println("user pass", mongo_user, mongo_pass)
		err = session.DB(database_server).Login(mongo_user, mongo_pass)
		if err != nil {
			//log.Fatalf(" %s", err)
			formatter.JSON(response, http.StatusInternalServerError, "Internal Server Error")
			return
		}
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