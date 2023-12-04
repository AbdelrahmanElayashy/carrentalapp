package operations

import (
	"rentalmanagement/logic/model"
	"time"
)

type CustomerOperationsInterface interface {
	RentCar(start time.Time, end time.Time, vin string, customerId string) (model.Rental, error)
	ListRentals(customerId string) ([]model.Rental, error)
	CancelRental(rentalId string) error
}
