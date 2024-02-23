package dto

import (
	"github.com/Sindhuja966/banking/errs"
)

type NewCustomerRequest struct {
	Id          int
	Name        string
	DateofBirth string
	City        string
	Zipcode     string
	Status      string
}

func (r NewCustomerRequest) Validate() *errs.AppError {
	if r.DateofBirth < "1800-01-01" {
		return errs.NewValidationError("To open a new customer you need to have a minimum age limit ")
	}

	return nil
}
