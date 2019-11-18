package main

import (
	"fmt"
	"log"
	//"os"
	"time"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "golang.org/x/crypto/bcrypt"
)

var mongo_server = "34.201.94.31"
var mongo_database = "paymentDB"
var mongo_collection = "paymentscollection"
var mongo_username = "admin"
var mongo_password = "password"






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


func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payments", addPayment(formatter)).Methods("POST")
	mx.HandleFunc("/payments", getAllPayments(formatter)).Methods("GET")
	mx.HandleFunc("/payment/{id}", getPayment(formatter)).Methods("GET")
	mx.HandleFunc("/payment/{id}", deletePayment(formatter)).Methods("DELETE")

}
//Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		  //   card := "12345678989427837"
    // hash, _ := HashCardDetails(card) // ignore error for the sake of simplicity

    // fmt.Println("card:", card)
    // fmt.Println("Hash:    ", hash)
		formatter.JSON(w, http.StatusOK, struct{ Test string }{" Payments API version 1.0 alive!"})
	}
}

func HashCardDetails(card string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(card), 14)
    return string(bytes), err
}

func addPayment(formatter *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongo_server)
		err := session.DB("admin").Login(mongo_username, mongo_password)
			if err != nil {
			formatter.JSON(writer, http.StatusInternalServerError, "Mongo Connection Error ")
			log.Fatal(err)
			return
			 panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongo_database).C(mongo_collection)
		fmt.Println(req.Body)

		var payment Payment
		 _ = json.NewDecoder(req.Body).Decode(&payment)

		uuid := uuid.NewV4()
		payment.PaymentID = uuid.String()
		t := time.Now()
		payment.PaymentDate = t.Format("2006-01-02 15:04:05")
		payment.OrderStatus = true
		 var card = payment.CardDetails
		 //fmt.Println("card:", card)
		 hash, _ := HashCardDetails(card)
		 payment.CardDetails = hash
		err = c.Insert(payment)
		if err != nil {
			formatter.JSON(writer, http.StatusNotFound, "Create Payment Error")
			return
		}
		fmt.Println("Create new payment:", payment)
		formatter.JSON(writer, http.StatusOK, payment)
	}
}




func getAllPayments(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongo_server)
		err := session.DB("admin").Login(mongo_username, mongo_password)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongo_database).C(mongo_collection)
		var result []bson.M
		err = c.Find(nil).All(&result)
		if err != nil {
			fmt.Println("error:" + err.Error())
			formatter.JSON(w, http.StatusNotFound, "Get All Payment Error")
			return
		}
		fmt.Println("getAllPaymentDetails:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

func getPayment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongo_server)
		err := session.DB("admin").Login(mongo_username, mongo_password)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongo_database).C(mongo_collection)
		var result bson.M
		params := mux.Vars(req)
		err = c.Find(bson.M{"orderid": params["id"]}).One(&result)
		fmt.Println("", err)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, "Get a Payment Error")
			return
		}
		fmt.Println("getAPaymentDetail:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

func deletePayment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, _ := mgo.Dial(mongo_server)
		err := session.DB("admin").Login(mongo_username, mongo_password)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Mongo Connection Error")
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongo_database).C(mongo_collection)
		var result Payment
		params := mux.Vars(req)
		err = c.Find(bson.M{"paymentid": params["id"]}).One(&result)
		fmt.Println("", err)
		if err != nil {
			fmt.Println("error:" + err.Error())
			formatter.JSON(w, http.StatusNotFound, "Delete a Payment Error")
			return
		} else {
			err = c.Remove(bson.M{"paymentid": result.PaymentID})
			if err != nil {
				fmt.Println("error:" + err.Error())
				formatter.JSON(w, http.StatusNotFound, "Delete Payment: deletion Error")
				return
			}
		}
		fmt.Println("DeletePaymentDetails:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}




