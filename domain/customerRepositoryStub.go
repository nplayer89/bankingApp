package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositorySub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "ChronicSmoke", City: "Fort Collins", Zipcode: "80525", DateOfBirth: "1989-01-01", Status: "1"},
		{Id: "1002", Name: "JoeDank", City: "Fort Collins", Zipcode: "80525", DateOfBirth: "1989-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
