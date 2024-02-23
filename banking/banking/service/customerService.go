package service

import (
	"context"

	"github.com/Sindhuja966/banking/domain"
	"github.com/Sindhuja966/banking/dto"
	"github.com/Sindhuja966/banking/errs"
)

//go:generate mockgen -destination=C:/Users/siboddap/banking/mocks/service/mockCustomerService.go -package=service github.com/Sindhuja966/banking/service customerService
type CustomerService interface {
	GetAllCustomer(ctx context.Context, status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id int) (*dto.CustomerResponse, *errs.AppError)
	CreateCustomer(request *dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError)
	SaveCustomer(customer domain.Customer) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetRepository() domain.CustomerRepository {
	return s.repo
}

func (s DefaultCustomerService) GetAllCustomer(ctx context.Context, status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(ctx, status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomer(id int) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) CreateCustomer(req *dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Customer{
		Id:          req.Id,
		Name:        req.Name,
		City:        req.City,
		Zipcode:     req.Zipcode,
		DateofBirth: req.DateofBirth,
		Status:      req.Status,
	}

	newCustomer, err := s.repo.SaveCustomer(a)
	if err != nil {
		return nil, err
	}
	response := newCustomer.ToNewCustomerResponseDto()
	return &response, nil
}

func (s DefaultCustomerService) SaveCustomer(customer domain.Customer) (*dto.CustomerResponse, *errs.AppError) {

	c, err := s.repo.SaveCustomer(customer)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil

}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
