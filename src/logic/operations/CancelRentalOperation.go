package operations

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (ops CustomerOperations) CancelRental(rentalId string) error {
	var msg string

	// Check if the rental exists
	exists, err := ops.rentalRepository.DoesRentalExist(rentalId)
	if err != nil {
		msg = fmt.Sprintf("Failed to check existence of rental with ID %s", rentalId)
		log.Error(msg, err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	if !exists {
		msg = fmt.Sprintf("Rental with ID %s does not exist", rentalId)
		log.Error(msg, err)
		return fmt.Errorf(msg)
	}

	// Delete rental
	if err := ops.rentalRepository.DeleteRental(rentalId); err != nil {
		msg = fmt.Sprintf("Failed to cancel the rental with ID %s", rentalId)
		log.Error(msg, err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}
