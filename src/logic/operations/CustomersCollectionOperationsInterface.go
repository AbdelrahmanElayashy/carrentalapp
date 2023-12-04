package operations

import (
	model "rentalmanagement/logic/model"
)

type CustomersCollectionOperationsInterface interface {
	RegisterCustomer(name string) (model.Customer, error)
	DeregisterCustomer(customerId string) error
}
