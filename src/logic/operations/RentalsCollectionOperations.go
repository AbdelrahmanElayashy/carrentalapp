package operations

import (
	"rentalmanagement/logic/model"
)

type RentalsCollectionOperations struct {
	rentalRepository model.RentalRepositoryInterface
	carRepository    model.CarRepositoryInterface
}

func NewRentalsCollectionOperations(rentalRepository model.RentalRepositoryInterface, carRepository model.CarRepositoryInterface) RentalsCollectionOperations {
	return RentalsCollectionOperations{rentalRepository: rentalRepository, carRepository: carRepository}
}
