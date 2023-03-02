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
