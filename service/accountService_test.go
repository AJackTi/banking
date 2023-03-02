package service

import (
	"testing"
	"time"

	realdomain "github.com/AJackTi/banking/domain"
	"github.com/AJackTi/banking/dto"
	"github.com/AJackTi/banking/errs"
	"github.com/AJackTi/banking/mocks/domain"
	"github.com/golang/mock/gomock"
)

func Test_NewAccount_Validate(t *testing.T) {
	// Arrange
	request := dto.NewAccountRequest{
		CustomerID:  "100",
		AccountType: "saving",
		Amount:      0,
	}
	service := NewAccountService(nil)

	// Act
	_, appErr := service.NewAccount(request)

	// Assert
	if appErr == nil {
		t.Error("failed while testing the new account validation")
	}
}

var (
	mockRepo *domain.MockAccountRepository
	service  AccountService
)

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_NewAccount_Error(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerID:  "100",
		AccountType: "saving",
		Amount:      6000,
	}

	account := &realdomain.Account{
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	// Act
	_, appErr := service.NewAccount(req)

	// Assert
	if appErr == nil {
		t.Error("Test failed while validating error for new account")
	}
}

func Test_NewAccount_Ok(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerID:  "100",
		AccountType: "saving",
		Amount:      6000,
	}

	account := &realdomain.Account{
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format(dbTSLayout),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	accountWithID := account
	mockRepo.EXPECT().Save(account).Return(accountWithID, nil)

	// Act
	newAccount, appErr := service.NewAccount(req)

	// Assert
	if appErr != nil {
		t.Error("Test failed while creating new account")
	}

	if newAccount.AccountID != accountWithID.AccountID {
		t.Error("Failed while matching new account id")
	}
}
