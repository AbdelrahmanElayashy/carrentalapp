package model

import (
	"time"
)

// The model of the car that is used by all internal operations
// This corresponds to the API Diagram
type Rental struct {
	RentalId   string
	StartDate  time.Time
	EndDate    time.Time
	CustomerId string
	Vin        Vin
}
