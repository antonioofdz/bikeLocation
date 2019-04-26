package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type UserDBCredentials struct {
	Username string
	Password string
}

type UserDBToken struct {
	Token string
}

type UserDB struct {
	UserDBToken
	Id      int
	Name    string
	Surname string
	Email   string
}

func loadRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/user/login", func(w http.ResponseWriter, req *http.Request) {
		var userDB *UserDBCredentials
		_ = json.NewDecoder(req.Body).Decode(&userDB)
		/*http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		json.NewEncoder(w).Encode(userDB)*/

		data, err := getUserCredentials(userDB)
		if err != nil {
			fmt.Println("hola3", err)
		}
		json.NewEncoder(w).Encode(data)

	}).Methods("POST")

	getUser := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		data, err := getUserUserByToken(req.Header.Get("token"))
		if err != nil {
			fmt.Println("sin user", err)
		}

		json.NewEncoder(w).Encode(data)

	})
	router.Handle("/user", checkToken(getUser)).Methods("GET")
	/*
	  router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	  router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")*/
	log.Fatal(http.ListenAndServe(":5001", router))
}

func checkToken(next http.Handler) http.Handler {
	checkValidGuidFn := func(uuid string) bool {
		r := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}")
		return r.MatchString(uuid)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(checkValidGuidFn(r.Header.Get("token")))
		if !checkValidGuidFn(r.Header.Get("token")) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	loadRoutes()
}

func open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dev-user:dev-password@/ermakina")
	if err != nil {
		return db, err
	}

	return db, nil
}

func getUserCredentials(userCredentials *UserDBCredentials) (*UserDBToken, error) {
	sqlGetUserCredentials := "SELECT users.token FROM users inner join users_credentials on users_credentials.id_user = users.id WHERE users_credentials.user=? and users_credentials.password=?"
	db, err := open()
	if err != nil {
		fmt.Println("hola2", err)
	}
	defer db.Close()

	var userDB UserDBToken
	err = db.QueryRow(sqlGetUserCredentials, &userCredentials.Username, &userCredentials.Password).Scan(&userDB.Token)
	if err != nil {
		fmt.Println("hola1", err)
	}

	return &userDB, nil
}

func getUserUserByToken(token string) (*UserDB, error) {
	sqlGetUserCredentials := "SELECT users.id, users.name, users.surname, users.email, users.token FROM users WHERE token=?"
	db, err := open()
	if err != nil {
		fmt.Println("hola2", err)
	}
	defer db.Close()

	var userDB UserDB
	err = db.QueryRow(sqlGetUserCredentials, token).Scan(&userDB.Id, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Token)
	if err != nil {
		fmt.Println("hola1", err)
	}

	return &userDB, nil
}
