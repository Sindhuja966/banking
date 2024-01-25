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
		//   {Id:"1001",Name:"Sindhu",City:"hyd",Zipcode:"50876",DateofBirth:"2000-02-2",Status:"1"},
		//   {Id:"1002",Name:"Arun",City:"hyd",Zipcode:"50876",DateofBirth:"2004-02-2",Status:"1"},
		{"1001", "Nokia", "hyd", "34234", "2000-03-03", "1"},
		{"1001", "sindhu", "hyd", "34234", "2004-03-02", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
