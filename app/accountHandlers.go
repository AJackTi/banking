package app

import (
	"encoding/json"
	"net/http"

	"github.com/AJackTi/banking/dto"
	"github.com/AJackTi/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	request.CustomerID = customerID

	account, appErr := ah.service.NewAccount(request)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.Message)
		return
	}

	writeResponse(w, http.StatusCreated, account)
}

func (ah *AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountID := vars["account_id"]
	customerID := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// build the request object
	request.AccountID = accountID
	request.CustomerID = customerID

	// make transaction
	account, appErr := ah.service.MakeTransaction(request)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, account)
}
