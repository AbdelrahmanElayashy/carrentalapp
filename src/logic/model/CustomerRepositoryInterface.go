package model

type CustomerRepositoryInterface interface {
	AddCustomer(customer Customer) (Customer, error)
	DeleteCustomer(customer Customer) error
	GetCustomer(customerId string) (Customer, error)
}
