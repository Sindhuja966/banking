package service

import (
	"testing"

	"github.com/Sindhuja966/banking/domain"
)

func TestNewCustomerService(t *testing.T) {
	repo := domain.NewCustomerRepositoryDb()
	service := NewCustomerService(repo)

	if service.GetRepository() == nil {
		t.Error("Expected a non repository in the service")
	}

}
