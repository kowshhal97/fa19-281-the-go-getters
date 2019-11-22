package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"os"
	"time"
	"log"

	

	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//uuid "github.com/satori/go.uuid"
	"github.com/unrolled/render"
	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)
/*
var mongodb_server = os.Getenv("AWS_MONGODB")
var mongodb_database = os.Getenv("MONGODB_DBNAME")
var mongodb_collection = os.Getenv("MONGODB_COLLECTION")
var mongodb_username = os.Getenv("MONGODB_USERNAME")
var mongodb_password = os.Getenv("MONGODB_PASSWORD")
*/
//hardcode for testing

//var mongodb_server = "mongodb://127.0.0.1:27017"
var mongodb_server = "mongodb://10.0.1.135:27017"
var mongodb_database = "cmpe281"
var mongodb_collection = "NewPizza"
var mongodb_username = "admin"
var mongodb_password = "admin"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(mx))
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/reviews/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/postReview", postReviewHandler(formatter)).Methods("POST")
	mx.HandleFunc("/getReviews/{itemName}", getReviewsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/deleteReview", deleteReviewHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/updateReview", updateReviewHandler(formatter)).Methods("PUT")

}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"The Go Getters Review API version 1.0 ALIVE!"})
	}
}

// API Post a Review Handler.
func postReviewHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var review Review
		_ = json.NewDecoder(req.Body).Decode(&review)
		fmt.Println("Review is: ",review.ItemName,  " ", review.Reviews)
		session, _ := mongo.Dial(mongodb_server)
		err := session.DB("admin").Login(mongodb_username, mongodb_password)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error ")
			return
		}
		defer session.Close()
		session.SetMode(mongo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		fmt.Println(req.Body)
		t := time.Now()
		entry := Review{
			ItemName: review.ItemName,
			Reviews: review.Reviews,
			ReviewDate : t.Format("2006-01-02 15:04:05"),
		}
		err = c.Insert(entry)
		if err != nil {
			fmt.Println("Error while adding reviews: ", err)
			formatter.JSON(w, http.StatusInternalServerError, struct{ Response error }{err})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Response string }{"Review added"})
		}
	}
}

//Get All Reviews for a Menu Item
func getReviewsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
	
		session, _ := mongo.Dial(mongodb_server)
		err := session.DB("admin").Login(mongodb_username, mongodb_password)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error ")
			return
		}
		defer session.Close()
		session.SetMode(mongo.Monotonic, true)

		params := mux.Vars(req)
		//var ItemName string = params["menuitemname"]
		//fmt.Println( "Item Name: ", ItemName )
		var ItemName string = params["itemName"]
		fmt.Println( "Item Name: ", ItemName )

		c := session.DB(mongodb_database).C(mongodb_collection)
		var results []Review
		err = c.Find(bson.M{"itemname": ItemName}).All(&results)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
		if len(results) > 0 {
			formatter.JSON(w, http.StatusOK, results)
		}else{
			formatter.JSON(w, http.StatusNoContent, struct{ Response string }{"No reviews found for the given ID"})
		}
	}
}

//Delete Last Review
func deleteReviewHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var m Review
		_ = json.NewDecoder(req.Body).Decode(&m)
		fmt.Println("Review is: ", m.Reviews)
		session, err := mongo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Reviews API (Delete) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mongo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		query := bson.M{
			"itemname": m.ItemName,
		}
		change := bson.M{"$pull": bson.M{ "reviews" :  bson.M{"$in": m.Reviews } }}
		err = c.Update(query, change)

		if err != nil {
			fmt.Println("Error while deleting reviews: ", err)
			formatter.JSON(w, http.StatusInternalServerError, struct{ Response error }{err})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Response string }{"Review deleted"})
		}
	}
}

func updateReviewHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var m Review
		_ = json.NewDecoder(req.Body).Decode(&m)
		fmt.Println("Review is: ", m.ItemName, " " , "Reviews", m.Reviews)
		session, err := mongo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Reviews API (Update) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mongo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		query := bson.M{
			"itemname": m.ItemName,
		}
		change := bson.M{"$push": bson.M{ "reviews" : bson.M{"$each": m.Reviews }}}
		err = c.Update(query, change)

		if err != nil {
			fmt.Println("Error while updating reviews: ", err)
			formatter.JSON(w, http.StatusInternalServerError, struct{ Response error }{err})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Response string }{"Review updated"})
		}
	}
}

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Vary", "Accept-Encoding")
}