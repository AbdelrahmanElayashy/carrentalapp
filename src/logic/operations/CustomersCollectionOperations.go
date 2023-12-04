package operations

import (
	"rentalmanagement/logic/model"
)

type CustomersCollectionOperations struct {
	customerRepository model.CustomerRepositoryInterface
}

func NewCustomersCollectionOperations(repository model.CustomerRepositoryInterface) CustomersCollectionOperations {
	return CustomersCollectionOperations{customerRepository: repository}
}
