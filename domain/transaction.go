package domain

import "github.com/AJackTi/banking/dto"

type Transaction struct {
	AccountID       string
	Amount          float64
	TransactionType string
	TransactionDate string
	TransactionID   string
}

func (t Transaction) IsWithdrawal() bool {
	return false
}

func (t Transaction) ToDto() *dto.TransactionResponse {
	return &dto.TransactionResponse{
		TransactionID:   t.TransactionID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
