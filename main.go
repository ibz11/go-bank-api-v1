package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ibz11/go-bank-api-v1.git/config"
	"github.com/ibz11/go-bank-api-v1.git/handlers"
)

func initializeRouter() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	r := mux.NewRouter()

	route := r.PathPrefix("/api/v1").Subrouter()
	//Routes
	route.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateAccount(w, r, db)
	}).Methods("POST")

	route.HandleFunc("/account/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAccount(w, r, db)
	}).Methods("GET")

	fmt.Println("Starting server...")
	fmt.Println("Go to http://localhost:4000")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func main() {
	initializeRouter()

}
