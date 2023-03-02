package service

import (
	"time"

	"github.com/AJackTi/banking/domain"
	"github.com/AJackTi/banking/dto"
	"github.com/AJackTi/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	newAccount, err := s.repo.Save(&domain.Account{
		AccountID:   "",
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	})
	if err != nil {
		return nil, err
	}

	return newAccount.ToNewAccountResponseDto(), nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
