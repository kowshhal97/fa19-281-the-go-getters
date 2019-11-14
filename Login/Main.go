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
	pgUrl, err := pq.ParseURL("postgres://gykxtcgv:SfJ_HK3OcQRaGpmoQjr1FAm2yaBqiPLi@salt.db.elephantsql.com:5432/gykxtcgv")
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
func responceJSON(w http.ResponseWriter,data interface{}){
	json.NewEncoder(w).Encode(data)

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
	hash,err:=bcrypt.GenerateFromPassword([]byte((user.Password)),10)
	if(err!=nil){
		log.Fatal()
	}
	user.Password=string(hash)
	query:="insert into users (email,password) values($1,$2) RETURNING id;"
	err=db.QueryRow(query,user.Email,user.Password).Scan(&user.ID)
	if(err!=nil){
		fmt.Println(error.message)
		error.message="Server error"
		respondwitherror(w,r,error)
	}
	user.Password=""
	w.Header().Set("Content-Type","application/json")
	responceJSON(w,user)
}
