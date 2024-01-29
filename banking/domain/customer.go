package domain

import (
	"github.com/Sindhuja966/banking/dto"
	"github.com/Sindhuja966/banking/errs"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.Status,
	}
}

func (a Customer) ToNewCustomerResponseDto() dto.NewCustomerResponse {
	return dto.NewCustomerResponse{Id: a.Id}
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError) //we are passing as a pointer bcz if in case id is passes as nil no customer is available we can do only with pointer
	SaveCustomer(Customer) (*Customer, *errs.AppError)
}
