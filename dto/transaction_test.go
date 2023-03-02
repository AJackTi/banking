package dto

import (
	"net/http"
	"testing"
)

func Test_TransactionTypeDepositWithdrawal(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		TransactionType: "invalid transaction type",
	}

	// Act
	appErr := request.Validate()

	// Assert
	if appErr.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error("Invalid message while testing transaction type")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func Test_Validate_Amount(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		TransactionType: DEPOSIT,
		Amount:          -100,
	}

	// Act
	appErr := request.Validate()

	// Assert
	if appErr.Message != "Amount cannot be less than zero" {
		t.Error("Invalid message while validating amount")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}
