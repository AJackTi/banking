package domain

import "github.com/AJackTi/banking/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}

// type CustomerRepositoryStub struct {
// 	customers []Customer
// }

// func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
// 	return s.customers, nil
// }

// func NewCustomerRepositoryStub() CustomerRepositoryStub {
// 	customers := []Customer{
// 		{"1001", "Ti", "Ho Chi Minh City", "100100", "25/04/1997", "1"},
// 		{"1002", "Ti 1", "Ho Chi Minh City", "100101", "26/04/1997", "0"},
// 	}

// 	return CustomerRepositoryStub{customers}
// }
