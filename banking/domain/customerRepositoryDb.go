package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/Sindhuja966/banking/errs"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	query := "select customer_id,name,city,zipcode,date_of_birth,status from customers"

	var rows *sql.Rows
	var err error

	if status != "" {
		query += " WHERE status = $1"
		rows, err = d.client.Query(query, status)
	} else {
		rows, err = d.client.Query(query)
	}

	if err != nil {
		log.Println("Error while querying customers table: " + err.Error())
		return customers, nil
	}
	defer rows.Close()

	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status); err != nil {
			log.Println("Error while scanning customers: " + err.Error())
			return customers, nil
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id int) (*Customer, *errs.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=$1"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) SaveCustomer(customer Customer) (*Customer, *errs.AppError) {

	log.Printf("Customer before insertion, %+v", customer)

	customerSql := `INSERT INTO customers(customer_id, name, city, zipcode, date_of_birth, status) VALUES ($1, $2, $3, $4, $5, $6)        
 RETURNING customer_id, name, city, zipcode, date_of_birth, status ` // Execute the prepared statement with the provided values

	row := d.client.QueryRow(customerSql, customer.Id, customer.Name, customer.City, customer.Zipcode, customer.DateofBirth, customer.Status)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	log.Printf("Order after insertion: %+v", customer)
	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dsn := "user=postgres password=Postgres dbname=postgres sslmode=disable"
	client, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
