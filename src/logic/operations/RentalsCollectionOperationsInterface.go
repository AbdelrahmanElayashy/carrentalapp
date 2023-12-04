package operations

import (
	model "rentalmanagement/logic/model"
	"time"
)

type RentalsCollectionOperationsInterface interface {
	ListAvailableCars(startDate, endDate time.Time) ([]model.Car, error)
}
