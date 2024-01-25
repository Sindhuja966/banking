package domain

import (
	"testing"
)

func TestNewCustomerRepositoryDb(t *testing.T) {
	repo := NewCustomerRepositoryDb()

	if repo.client == nil {
		t.Error("Expected a no client in the repository")
	}

}
