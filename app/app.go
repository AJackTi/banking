package app

import (
	"log"
	"net/http"

	"github.com/AJackTi/banking/domain"
	"github.com/AJackTi/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	// wiring
	domain := domain.NewCustomerRepositoryDb()
	service := service.NewCustomerService(domain)
	ch := CustomerHandlers{service}

	mux := mux.NewRouter()

	// define routes
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
