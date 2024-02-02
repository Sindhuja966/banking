package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// used to create new dummy customers
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1001, "Nokia", "hyd", "34234", "2000-03-03", "1"},
		{1001, "sindhu", "hyd", "34234", "2004-03-02", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
