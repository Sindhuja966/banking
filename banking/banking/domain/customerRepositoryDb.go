package domain

import (
	"context"
	"log"
	"time"

	"github.com/Sindhuja966/banking/errs"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepositoryDb struct {
	client *gorm.DB
}

func (d CustomerRepositoryDb) FindAll(ctx context.Context, status string) ([]Customer, *errs.AppError) {
	time.Sleep(5 * time.Second)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, errs.NewUnexpectedError("Deadline exceeded")
	}
	var customers []Customer

	if status != "" {
		result := d.client.Where("status = ?", status).Find(&customers)
		if result.Error != nil {
			log.Println("Error while querying customers table: " + result.Error.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		}
	} else {
		result := d.client.Find(&customers)
		if result.Error != nil {
			log.Println("Error while querying customers table: " + result.Error.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id int) (*Customer, *errs.AppError) {
	var customer Customer
	result := d.client.First(&customer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while querying customer: " + result.Error.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &customer, nil
}

func (d CustomerRepositoryDb) SaveCustomer(customer Customer) (*Customer, *errs.AppError) {
	log.Printf("Customer before insertion: %+v", customer)
	result := d.client.Create(&customer)
	if result.Error != nil {
		log.Println("Error while inserting customer: " + result.Error.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	log.Printf("Customer after insertion: %+v", customer)
	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dsn := "user=postgres password=Postgres dbname=postgres sslmode=disable"
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return CustomerRepositoryDb{client}
}
