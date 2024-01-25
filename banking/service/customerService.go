package service

import (
	"github.com/Sindhuja966/banking/domain"
	"github.com/Sindhuja966/banking/dto"
	"github.com/Sindhuja966/banking/errs"
)

//go:generate mockgen -destination=C:/Users/siboddap/banking/mocks/service/mockCustomerService.go -package=service github.com/Sindhuja966/banking/service customerService
type CustomerService interface {
	//GetRepository() domain.CustomerRepository
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	CreateCustomer(dto.NewCustomerRequest) (*dto.CustomerResponse, *errs.AppError)
	SaveCustomer(Customer domain.Customer) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetRepository() domain.CustomerRepository {
	return s.repo
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) CreateCustomer(customer domain.Customer) (*dto.CustomerResponse, *errs.AppError) {
	a := domain.Customer{
		Id:          "",
		Name:        "",
		City:        "",
		Zipcode:     "",
		DateofBirth: "",
		Status:      "",
	}

	return s.repo.SaveCustomer(a)
}
func (s DefaultCustomerService) SaveCustomer(customer domain.Customer) (*dto.CustomerResponse, *errs.AppError) {
	return s.repo.SaveCustomer(customer)

}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

// func NewMockCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
// 	return DefaultCustomerService{repo: repository}
// }
