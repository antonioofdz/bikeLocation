package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/antonioofdz/personalProjectDra/pkg/models"
)

func LoadRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/login", loginUserController).Methods("POST")
	router.HandleFunc("/signin", signInController).Methods("POST")
	router.Handle("/user", CheckToken(http.HandlerFunc(getUserByTokenController))).Methods("GET")
	router.Handle("/bikes", CheckToken(http.HandlerFunc(getListBikesController))).Methods("GET")
	router.Handle("/bikes/book", CheckToken(http.HandlerFunc(bookBikeController))).Methods("POST")
	router.Handle("/bikes/return", CheckToken(http.HandlerFunc(returnBikeController))).Methods("POST")

	log.Fatal(http.ListenAndServe(":5001", router))
}

// Controlador para logear a un usuario
func loginUserController(w http.ResponseWriter, req *http.Request) {
	var userDB *models.UserDBCredentials
	if err := json.NewDecoder(req.Body).Decode(&userDB); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
    	w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	data, err := getUserCredentials(userDB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
    	w.Write([]byte("Something bad happened! [/user/login]"))
	}
	json.NewEncoder(w).Encode(data)
}

// Controlador para dar de alta a un nuevo usuario
func signInController(w http.ResponseWriter, req *http.Request) {
	var signInUserDB *models.SignInUserDB
	if err := json.NewDecoder(req.Body).Decode(&signInUserDB); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
    	w.Write([]byte("Error Parsing BODY! [/signin]"))
	}
	
	err := signInUser(signInUserDB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
    	w.Write([]byte("Something bad happened! [/signin] \n"+ err.Error()))
	}
	json.NewEncoder(w).Encode(http.StatusOK)
}

// Controlador que obtiene un usuario por su Token
func getUserByTokenController(w http.ResponseWriter, req *http.Request) {
	data, err := getUserUserByToken(req.Header.Get("token"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Something bad happened! [/user]"))
	}

	json.NewEncoder(w).Encode(data)
}

// Controlador para obtener un listado de bicicletas
func getListBikesController(w http.ResponseWriter, req *http.Request) {
	data, err := getListBikes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Something bad happened! [/bikes]"))
	}

	json.NewEncoder(w).Encode(data)
}

// Controlador para reservar una bicicleta
func bookBikeController(w http.ResponseWriter, req *http.Request) {
	var bookBikeModel models.BookBike
	err := json.NewDecoder(req.Body).Decode(&bookBikeModel)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
    	w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	if err = bookBike(&bookBikeModel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Something bad happened! [/bikes/book]"))
	}

	json.NewEncoder(w).Encode(http.StatusOK)
}

// Controlador para devolver una bici ya alquilada
func returnBikeController(w http.ResponseWriter, req *http.Request) {
	var bookBikeModel models.BookBike
	err := json.NewDecoder(req.Body).Decode(&bookBikeModel)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
    	w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	if err = endBookBike(&bookBikeModel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Something bad happened! [/bikes/return]"))
	}

	json.NewEncoder(w).Encode(http.StatusOK)
}