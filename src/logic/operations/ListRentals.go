package operations

import (
	"fmt"
	"rentalmanagement/logic/model"

	log "github.com/sirupsen/logrus"
)

func (ops CustomerOperations) ListRentals(customerId string) ([]model.Rental, error) {
	var msg string

	// // Check if the customer exists
	// exists, err := ops.customerRepository.GetCustomer(customerId)
	// if err != nil {
	// 	msg = fmt.Sprintf("Failed to check existence of customer with ID %s", customerId)
	// 	return nil, fmt.Errorf("%s: %w", msg, err)
	// }

	// if !exists {
	// 	msg = fmt.Sprintf("Customer with ID %s does not exist", customerId)
	// 	return nil, fmt.Errorf(msg)
	// }

	// List all rentals by customer ID
	rentals, err := ops.rentalRepository.ListRentalsByCustomerId(customerId)
	if err != nil {
		msg = fmt.Sprintf("Failed to list rentals for customer with ID %s", customerId)
		log.Error(msg, err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return rentals, nil
}
