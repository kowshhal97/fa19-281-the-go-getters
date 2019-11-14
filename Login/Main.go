package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"log"
	"net/http"
"golang.org/x/crypto/bcrypt"
)

type User struct{
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type Error  struct{
	message string `"message"`
}
type JWT struct{
	Token string `json:"token"`
}
var db *sql.DB
func main() {
	pgUrl, err := pq.ParseURL("postgres://gykxtcgv:SfJ_HK...@salt.db.elephantsql.com:5432/gykxtcgv ")
	if(err!=nil){
		log.Fatal()
	}
	db,err=sql.Open("postgres",pgUrl)
	if(err!=nil){
		log.Fatal()
	}
	err=db.Ping()
router:=mux.NewRouter()
router.HandleFunc("/signup",signupPost).Methods("POST")

http.ListenAndServe(":8080",router)
}
func respondwitherror(w http.ResponseWriter, r* http.Request,error Error){
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(error)
}
func responceJSON(w http.ResponseWriter,r* http.Request,data interface{}){

}
func signupPost(w http.ResponseWriter,r *http.Request){
	var user User
	var error Error
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if(user.Email==""){
		error.message="email is missing"
		respondwitherror(w,r,error)
		return
	}
	if(user.Password=="") {
		error.message = "password is missing"
		respondwitherror(w, r, error)
		return
	}
}
func indexhandler(w http.ResponseWriter, r* http.Response){
	w.Write([]byte("this is index page"))
}
