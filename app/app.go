package app

import (
	"github.com/ajackti/banking/domain"
	"github.com/ajackti/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
