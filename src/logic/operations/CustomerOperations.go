package operations

import (
	"rentalmanagement/logic/model"
)

type CustomerOperations struct {
	rentalRepository model.RentalRepositoryInterface
	carRepository    model.CarRepositoryInterface
}

func NewCustomerOperations(rentalRepository model.RentalRepositoryInterface, carRepository model.CarRepositoryInterface) CustomerOperations {
	return CustomerOperations{rentalRepository: rentalRepository, carRepository: carRepository}
}
