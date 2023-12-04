package operations

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"
	"rentalmanagement/logic/model"
)

func (ops CustomersCollectionOperations) DeregisterCustomer(customerID string) error {
	var msg string

	// Check if the customer exists
	customer, err := ops.customerRepository.GetCustomer(customerID)
	if err != nil {
		msg = fmt.Sprintf("Failed to check existence of customer with ID %s", customerID)
		log.Error(msg, err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	if cmp.Equal(customer, model.Customer{}) {
		msg = fmt.Sprintf("customer with ID %s does not exist", customerID)
		log.Error(msg, err)
		return fmt.Errorf(msg)
	}

	// Delete customer
	if err := ops.customerRepository.DeleteCustomer(customer); err != nil {
		msg = fmt.Sprintf("Failed to deregister the customer with ID %s", customerID)
		log.Error(msg, err.Error())
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}
