package dto

import (
	"github.com/AJackTi/banking/errs"
)

type TransactionRequest struct {
	AccountID       string
	CustomerID      string
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

const (
	WITHDRAWAL = "withdraw"
	DEPOSIT    = "deposit"
)

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil
}

type TransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
