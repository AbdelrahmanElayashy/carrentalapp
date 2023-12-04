package operations

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"rentalmanagement/logic/model"
)

func (ops CustomersCollectionOperations) RegisterCustomer(name string) (model.Customer, error) {
	var msg string

	customer, err := ops.customerRepository.AddCustomer(model.Customer{Name: name})
	if err != nil {
		msg = fmt.Sprintf("Failed to add the customer with Name %s", customer.Name)
		log.Error(msg, err)
		return model.Customer{}, fmt.Errorf("%s: %w", msg, err)
	}

	return customer, nil
}
