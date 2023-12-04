package operations

import (
	"fmt"
	"rentalmanagement/logic/model"
	"time"

	log "github.com/sirupsen/logrus"
)

func (ops CustomerOperations) RentCar(start time.Time, end time.Time, vin string, customerId string) (model.Rental, error) {
	var msg string

	// Validate if StartDate is before EndDate
	if start.After(end) || start.Equal(end) {
		msg = "StartDate must be before EndDate"
		log.Warn(msg)
		return model.Rental{}, fmt.Errorf(msg)
	}

	// Check if the car exists in the car repository
	carExists, err := ops.carRepository.CarExists(model.Vin{Vin: vin})
	if err != nil {
		msg = "error checking car existence"
		log.Error(msg, err)
		return model.Rental{}, fmt.Errorf("%s: %w", msg, err)
	}

	if !carExists {
		msg = fmt.Sprintf("car with VIN %s does not exist", vin)
		log.Warn(msg)
		return model.Rental{}, fmt.Errorf(msg)
	}

	// Check if the car is available for the specified time range
	isCarAvailable, err := ops.rentalRepository.IsCarAvailableForRental(vin, start, end)
	if err != nil {
		msg = "error checking car availability"
		log.Error(msg, err)
		return model.Rental{}, fmt.Errorf("%s: %w", msg, err)
	}

	if !isCarAvailable {
		msg = fmt.Sprintf("car with VIN %s is not available for the specified time range", vin)
		log.Warn(msg)
		return model.Rental{}, fmt.Errorf(msg)
	}

	rental := model.Rental{
		StartDate:  start,
		EndDate:    end,
		CustomerId: customerId,
		Vin:        model.Vin{Vin: vin},
	}
	// Add the rental
	rental, err = ops.rentalRepository.AddRental(rental)
	if err != nil {
		msg = "error adding rental"
		log.Error(msg, err)
		return model.Rental{}, fmt.Errorf("%s: %w", msg, err)
	}

	return rental, nil
}
