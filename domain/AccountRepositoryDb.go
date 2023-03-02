package domain

import (
	"fmt"
	"strconv"

	"github.com/AJackTi/banking/errs"
	"github.com/AJackTi/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a *Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES(?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error(fmt.Sprintf("Error while creating new account: %v", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("Error while getting last insert id for new account: %v", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountID = strconv.FormatInt(id, 10)
	return a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
