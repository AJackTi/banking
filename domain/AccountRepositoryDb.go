package domain

import (
	"fmt"
	"strconv"

	"github.com/AJackTi/banking-lib/errs"
	"github.com/AJackTi/banking-lib/logger"
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

// transaction = make an entry in the transaction table + update the balance in the accounts table
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error(fmt.Sprintf("Error while starting a new transaction for bank account transaction: %v", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`, t.AccountID, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? WHERE account_id = ?`, t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? WHERE account_id = ?`, t.Amount, t.AccountID)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error(fmt.Sprintf("Error while saving transaction: %v\n", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		logger.Error(fmt.Sprintf("Error while commiting transaction: %v\n", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// getting the last transaction ID from the transaction table
	transactionID, err := result.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("Error while getting the last transaction id: %v\n", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountID)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionID = strconv.FormatInt(transactionID, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountID string) (*Account, *errs.AppError) {
	var account Account
	findBySql := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?"
	if err := d.client.Get(&account, findBySql, accountID); err != nil {
		logger.Error(fmt.Sprintf("Error while querying accounts table: %v\n", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &account, nil
}
