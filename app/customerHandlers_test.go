package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AJackTi/banking-lib/errs"
	"github.com/AJackTi/banking/dto"
	"github.com/AJackTi/banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var (
	router      *mux.Router
	ch          CustomerHandlers
	mockService *service.MockCustomerService
)

func setup(t *testing.T) func() {
	// Arrange
	ctrl := gomock.NewController(t)

	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}

	router = mux.NewRouter()

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_GetCustomers_Ok(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	dummyService := []dto.CustomerResponse{
		{ID: "1001", Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{ID: "1002", Name: "Rob", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
	}

	mockService.EXPECT().GetAllCustomer("").Return(dummyService, nil)

	router.HandleFunc("/customers", ch.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_GetCustomers_Error(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some database error"))
	ch := CustomerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
