package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/unrolled/render"
	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = os.Getenv("AWS_MONGODB")
var mongodb_database = os.Getenv("MONGODB_DBNAME")
var mongodb_collection = os.Getenv("MONGODB_COLLECTION")
var mongodb_username = os.Getenv("MONGODB_USERNAME")
var mongodb_password = os.Getenv("MONGODB_PASSWORD")

//hardcode for testing
/*
var mongodb_server = "mongodb://127.0.0.1:27017"
var mongodb_database = "cmpe281"
var mongodb_collection = "pizza"
var mongodb_username = "admin"
var mongodb_password = "admin"
*/
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
	mx.HandleFunc("/reviews/{menuItem}", getAllReviewsforItem(formatter)).Methods("GET")
	mx.HandleFunc("/reviews",  addNewReview(formatter)).Methods("POST")
	//mx.HandleFunc("/payment/{id}", getPaymentDetailsOfOne(formatter)).Methods("GET")
	//mx.HandleFunc("/payment/{id}", deletePaymentDetailsOfOne(formatter)).Methods("DELETE")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"The Go Getters Review API version 1.0 ALIVE!"})
	}
}

// API Payments Handler
func addNewReview(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		session, _ := mongo.Dial(mongodb_server)
		err := session.DB("admin").Login(mongodb_username, mongodb_password)
		if err != nil {
			formatter.JSON(writer, http.StatusInternalServerError, "Mongo Connection Error ")
			return
		}
		defer session.Close()
		session.SetMode(mongo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)
		fmt.Println(req.Body)

		var review Review
		_ = json.NewDecoder(req.Body).Decode(&review)
		fmt.Printf("", review)

		uuid,_ := uuid.NewV4()
		review.ReviewId = uuid.String()
		t := time.Now()
		review.ReviewDate = t.Format("2006-01-02 15:04:05")
		
		err = collection.Insert(review)
		if err != nil {
			formatter.JSON(writer, http.StatusNotFound, "Create Review Error")
			return
		}
		fmt.Println("Create new review:", review)
		formatter.JSON(writer, http.StatusOK, review)
	}
}

//API to get all Payment
func getAllReviewsforItem(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mongo.Dial(mongodb_server)
		err := session.DB("admin").Login(mongodb_username, mongodb_password)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			return
		}
		defer session.Close()
		session.SetMode(mongo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		parameter := mux.Vars(req)
		fmt.Println("parameters",parameter)
		var menu_item string = parameter["menuItem"]
		fmt.Println("Input userId is: ", menu_item)
		var result []bson.M
		//fmt.Println(c.Find(bson.M{/*"MenuItemId": "pizza"*/}).One(&result))
		 err = c.Find(bson.M{"menuItem": "menu_item"}).One(&result)
		if err != nil {
			fmt.Println("error:" + err.Error())
			formatter.JSON(w, http.StatusNotFound, "Get All Review for menu a item Error")
			return
		}
		fmt.Println("getAllPaymentDetails:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}



/*

  	-- Gumball MongoDB Collection (Create Document) --

    db.gumball.insert(
	    { 
	      Id: 1,
	      CountGumballs: NumberInt(202),
	      ModelNumber: 'M102988',
	      SerialNumber: '1234998871109' 
	    }
	) ;

    -- Gumball MongoDB Collection - Find Gumball Document --

    db.gumball.find( { Id: 1 } ) ;

    {
        "_id" : ObjectId("54741c01fa0bd1f1cdf71312"),
        "Id" : 1,
        "CountGumballs" : 202,
        "ModelNumber" : "M102988",
        "SerialNumber" : "1234998871109"
    }

    -- Gumball MongoDB Collection - Update Gumball Document --

    db.gumball.update( 
        { Dd: 1 }, 
        { $set : { CountGumballs : NumberInt(10) } },
        { multi : false } 
    )

    -- Gumball Delete Documents

    db.gumball.remove({})

 */

