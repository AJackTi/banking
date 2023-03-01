package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AJackTi/banking/domain"
	"github.com/AJackTi/banking/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

func Start() {
	sanityCheck()

	// wiring
	domain := domain.NewCustomerRepositoryDb()
	service := service.NewCustomerService(domain)
	ch := CustomerHandlers{service}

	mux := mux.NewRouter()

	// define routes
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}
