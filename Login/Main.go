package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Error  struct {
	Message string `json:"message"`
}
type JWT struct {
	Token string `json:"token"`
}
var db *sql.DB
func main() {
	pgUrl, err := pq.ParseURL("postgres://gykxtcgv:SfJ_HK3OcQRaGpmoQjr1FAm2yaBqiPLi@salt.db.elephantsql.com:5432/gykxtcgv")
	if (err != nil) {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", pgUrl)
	if (err != nil) {
		log.Fatal(err)
	}
	err = db.Ping()
	if(err!=nil){
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/signup", signupPost).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected",TokenVerifyMiddleware(protectedEndPoint)).Methods("GET")
	http.ListenAndServe(":8080", router)
}
func respondwitherror(w http.ResponseWriter,status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
func responceJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
func JSONDecoder(r*http.Request,user User) User {
	json.NewDecoder(r.Body).Decode(&user)
	return user
}
func GenerateToken(user User)(string,error) {
	var err error
	secret := "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
func login(w http.ResponseWriter, r* http.Request) {
	var user User
	var error Error
	var jwt JWT
	//validator code
	user = JSONDecoder(r, user)
	tokenString, err := GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}
	password := user.Password
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err = row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = ("user does not exist")
			respondwitherror(w,http.StatusBadRequest , error)
			return
		} else {
			log.Fatal(err)
		}
	}
	hashedPassword:=user.Password
	err=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	if(err!=nil){
		error.Message ="Invalid Password"
		respondwitherror(w,http.StatusUnauthorized,error)
		return
	}
	w.WriteHeader(http.StatusOK)
	jwt.Token=tokenString
	responceJSON(w,jwt.Token)
}
func signupPost(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error
	//json.NewDecoder(r.Body).Decode(&user)
	user = JSONDecoder(r, user)
	fmt.Println(user)
	//Validator code
	if (user.Email == "") {
		error.Message = "email is missing"
		respondwitherror(w, http.StatusBadRequest, error)
		return
	}
	if (user.Password == "") {
		error.Message = "password is missing"
		respondwitherror(w, http.StatusBadRequest, error)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte((user.Password)), 10)
	if (err != nil) {
		log.Fatal(err)
	}
	user.Password = string(hash)
	query := "insert into users (email,password) values($1,$2) RETURNING id;"
	err = db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	if (err != nil) {
		fmt.Println(error.Message)
		error.Message = "Server error"
		respondwitherror(w, http.StatusInternalServerError, error)
		return
	}
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	responceJSON(w, user)
}
func protectedEndPoint(w http.ResponseWriter, r* http.Request){
	fmt.Println("You have accessed protected endpoint")
}
func TokenVerifyMiddleware(next http.HandlerFunc)http.HandlerFunc{

	return http.HandlerFunc(func(w http.ResponseWriter,r* http.Request){
		var errorObject Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				respondwitherror(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				respondwitherror(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token."
			respondwitherror(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}

