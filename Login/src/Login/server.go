
package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var mongodb_server = "mongodb://admin:admin@10.0.1.18:27017"
var mongodb_database = "admin"
var mongodb_collection = "users"

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	corsObj := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.Use(corsObj)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/login", loginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/signup", signupHandler(formatter)).Methods("POST")
}
func respondwitherror(w http.ResponseWriter,status int, error Error) {
	w.WriteHeader(status)
	err:=json.NewEncoder(w).Encode(error)
	if(err!=nil){
		log.Fatal(err)
	}
}
func responceJSON(w http.ResponseWriter, data interface{}) {
	err:=json.NewEncoder(w).Encode(data)
	if(err!=nil){
		log.Fatal(err)
	}
}
func JSONDecoder(r*http.Request,user Users) Users {
	err:=json.NewDecoder(r.Body).Decode(&user)
	if(err!=nil){
		log.Fatal(err)
	}
	return user
}
func signupHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		var error Error;
		setDefaultHeaders(w)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		decoder := json.NewDecoder(req.Body)
		var user Users
		err1 := decoder.Decode(&user)
		if err1 != nil {
			panic(err1)
		}
		UserName := user.UserName
		Password := user.Password
		FirstName := user.FirstName
		LastName := user.LastName
		fmt.Println(UserName)
		fmt.Println(Password)
		user.Id,err =session.DB(mongodb_database).C(mongodb_collection).Count()
		if(err!=nil){
			log.Fatal(err)
		}
		user.Id++
		Id:=user.Id
		if (user.UserName == "") {
			error.Message = "username is missing"
			respondwitherror(w, http.StatusBadRequest, error)
			return
		}
		if (user.Password == "") {
			error.Message = "password is missing"
			respondwitherror(w, http.StatusBadRequest, error)
			return
		}
		result := Users{}
		err3 := c.Find(bson.M{"username": UserName}).One(&result)
		fmt.Println(err3)
		if (result.UserName != "")  {
			err:=formatter.JSON(w, http.StatusOK, "false")
			if(err!=nil) {
				log.Fatal(err)
			}
		}

		hash, err := bcrypt.GenerateFromPassword([]byte((user.Password)), 10)
		if (err != nil) {
			log.Fatal(err)
		}
		Password=string(hash);
		if (result.UserName == ""){

			err4 := c.Insert(&Users{Id:Id,UserName: UserName, Password: Password, FirstName: FirstName, LastName: LastName})

			if err4 != nil {
				panic(err4)
			}
			fmt.Println(err4)
			user.Password=""
			err:=formatter.JSON(w, http.StatusOK, user)
			if(err!=nil){
				log.Fatal(err)
			}
		}
	}

}

func loginHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		var error Error
		setDefaultHeaders(w)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		decodeddata := json.NewDecoder(req.Body)

		var user Users
		err1 := decodeddata.Decode(&user)
		if err1 != nil {
			panic(err1)
		}

		UserName := user.UserName
		Password := user.Password
		fmt.Println(UserName)
		fmt.Println(Password)

		result := Users{}
		err2 := c.Find(bson.M{"username": UserName}).One(&result)
		fmt.Println(err2)
		fmt.Println("result")
		fmt.Println(result)
		user.Id=result.Id
		hashedPassword:=result.Password
		err=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(Password))
		if(err!=nil){
			error.Message ="Invalid Password"
			respondwitherror(w,http.StatusUnauthorized,error)
			return
		} else{
			err:=formatter.JSON(w, http.StatusOK, result)
			if(err!=nil){
				log.Fatal(err)
			}
		}
	}

}

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err:=formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
		if(err!=nil){
			log.Fatal(err)
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
