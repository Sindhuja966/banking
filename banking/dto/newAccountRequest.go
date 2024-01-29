package dto

import (
	"github.com/Sindhuja966/banking/errs"
)

type NewCustomerRequest struct {
	Id          string
	Name        string
	DateofBirth string
	City        string
	Zipcode     string
	Status      string
}

func (r NewCustomerRequest) Validate() *errs.AppError {
	if r.Zipcode < "6000" {
		return errs.NewValidationError("To open a new customer you need to have a crct zipcode ")
	}

	return nil
}
