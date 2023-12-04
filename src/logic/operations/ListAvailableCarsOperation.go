package operations

import (
	"fmt"
	"math/rand"
	"rentalmanagement/logic/model"
	"time"

	log "github.com/sirupsen/logrus"
)

func (ops RentalsCollectionOperations) ListAvailableCars(startDate, endDate time.Time) ([]model.Car, error) {
	var msg string

	// Validate if StartDate is before EndDate
	if startDate.After(endDate) || startDate.Equal(endDate) {
		msg = "StartDate must be before EndDate"
		log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	// Fetch all cars; ideally, this would be optimized to only fetch available cars from database by querying
	cars, err := ops.carRepository.ListAllCars()
	if err != nil {
		msg = "Failed to list all cars"
		log.Error(msg, err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	// Pre-allocate the slice to avoid dynamic resizing
	availableCars := make([]model.Car, 0, len(cars))

	for _, car := range cars {
		isAvailable, err := ops.rentalRepository.IsCarAvailableForRental(car.Vin.Vin, startDate, endDate)
		if err != nil {
			msg = fmt.Sprintf("Failed to check availability for car with VIN %s", car.Vin.Vin)
			log.Error(msg, err)
			return nil, fmt.Errorf("%s: %w", msg, err)
		}

		// Set PricePerDay
		prices := []int{50, 100}
		availableCar := model.Car{
			Vin:         car.Vin,
			Brand:       car.Brand,
			Model:       car.Model,
			PricePerDay: prices[rand.Int()%len(prices)],
		}

		if isAvailable {
			availableCars = append(availableCars, availableCar)
		}
	}

	return availableCars, nil
}
