package domain

import (
	"github.com/AJackTi/banking/dto"
	"github.com/AJackTi/banking/errs"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountID: a.AccountID}
}

type AccountRepository interface {
	Save(*Account) (*Account, *errs.AppError)
}
