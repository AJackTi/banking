package domain

import (
	"time"

	"github.com/AJackTi/banking-lib/errs"
	"github.com/AJackTi/banking/dto"
)

const dbTSLayout = "2006-01-02T15:04:05Z07:00"

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

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/AJackTi/banking/domain AccountRepository
type AccountRepository interface {
	Save(*Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

func NewAccount(customerID, accountType string, amount float64) *Account {
	return &Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format(dbTSLayout),
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
