package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()

	// define routes
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
