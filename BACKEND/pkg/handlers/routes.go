package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/antonioofdz/personalprojectdra/pkg/models"
)

func LoadRoutes() {
	fmt.Println("Iniciando rutas..213213123sasadsda")
	router := mux.NewRouter()

	// LOGIN Y DEMAS PARA USUARIOS
	router.HandleFunc("/user/login", loginUserController).Methods("POST")
	router.HandleFunc("/user/add", signInController).Methods("PUT")

	//Alquilar y devolver
	router.HandleFunc("/bikes/book", bookBikeController).Methods("POST")
	router.HandleFunc("/bikes/return", returnBikeController).Methods("POST")

	//CRUD BICIS
	router.HandleFunc("/bikes/add", insertBikeController).Methods("PUT")
	router.HandleFunc("/bikes", updateBikeController).Methods("POST")
	router.HandleFunc("/bikes/{id}", getBikeController).Methods("GET")
	router.HandleFunc("/bikes/delete", deleteBikeController).Methods("POST")
	router.HandleFunc("/bikes", getListBikesController).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Fatal(http.ListenAndServe(":5002", c.Handler(router)))
}

func initDBController(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Error rutas.. initDBController")
	if err := InitDBHandler(); err != nil {
		fmt.Println("Error rutas.. ", err)

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error Parsing BODY! [/initDB]"))
	}
}

// Controlador para logear a un usuario
func loginUserController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

	var userDB *models.UserDBCredentials
	if err := json.NewDecoder(req.Body).Decode(&userDB); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	token, err := getUserCredentials(userDB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Something bad happened! [/user/login]"))
	}

	dataResult, err := getUserUserByToken(token.Token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/user]"))
	}

	json.NewEncoder(w).Encode(dataResult)
}

// Controlador para dar de alta a un nuevo usuario
func signInController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

	var signInUserDB *models.SignInUserDB
	if err := json.NewDecoder(req.Body).Decode(&signInUserDB); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error Parsing BODY! [/signin]"))
	}

	err := signInUser(signInUserDB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Something bad happened! [/signin] \n" + err.Error()))
	}
	json.NewEncoder(w).Encode(http.StatusOK)
}

// Controlador para obtener un listado de bicicletas
func getListBikesController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

	fmt.Println("Lllamando a getListBikesController")
	data, err := getListBikes()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/bikes]"))
	}

	json.NewEncoder(w).Encode(data)
}

// Controlador para obtener un listado de bicicletas
func deleteBikeController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

	var bikeID *models.BikeId
	if err := json.NewDecoder(req.Body).Decode(&bikeID); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	fmt.Println("Lllamando a deleteBikeController")
	err := DeleteBike(bikeID.Id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/bikes]"))
	}
}

// Controlador para obtener un listado de bicicletas
func updateBikeController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

	var bikeFull *models.BikeFull
	if err := json.NewDecoder(req.Body).Decode(&bikeFull); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	fmt.Println("Lllamando a updateBikeController")
	err := UpdateBike(bikeFull)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/bikes]"))
	}

}

// Controlador para obtener un listado de bicicletas
func insertBikeController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

	var bikeFull *models.BikeTest
	if err := json.NewDecoder(req.Body).Decode(&bikeFull); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	fmt.Println("Lllamando a insertBikeController")
	err := InsertBike(bikeFull)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/bikes]"))
	}

}

// Controlador para obtener un listado de bicicletas
func getBikeController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

	params := mux.Vars(req)
	idHeader := params["id"]
	fmt.Println("IDHEADER" + idHeader)
	if idHeader == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CANT FIND ID PARAMETER IN QUERY"))
	}

	id, err := strconv.ParseInt(idHeader, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID PARAMTER WITH BAD FORMAT"))
	}

	data, err := getBike(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/bike]"))
	}

	json.NewEncoder(w).Encode(data)
}

// Controlador para reservar una bicicleta
func bookBikeController(w http.ResponseWriter, req *http.Request) {
	allowRequestCors(w)

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

func allowRequestCors(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
