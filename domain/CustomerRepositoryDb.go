package domain

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AJackTi/banking/errs"
	"github.com/AJackTi/banking/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var (
		err  error
		rows *sql.Rows
	)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error(fmt.Sprintf("Error while querying customer table: %v\n", err.Error()))
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status); err != nil {
			logger.Error(fmt.Sprintf("Error while scanning customers: %v\n", err.Error()))
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer
	if err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error(fmt.Sprintf("Error while scanning customer: %v\n", err.Error()))
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:ductrong@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
