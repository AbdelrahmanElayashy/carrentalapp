package model

// The model of the car that is used by all internal operations
// This corresponds to the API Diagram
type Car struct {
	Vin         Vin
	Brand       string
	Model       string
	PricePerDay int
}
