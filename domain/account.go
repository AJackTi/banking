package domain

import (
	"github.com/AJackTi/banking/dto"
	"github.com/AJackTi/banking/errs"
)

type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountID: a.AccountID}
}

type AccountRepository interface {
	Save(*Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
