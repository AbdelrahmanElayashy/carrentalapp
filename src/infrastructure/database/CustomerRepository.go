package database

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"rentalmanagement/infrastructure/database/entities"
	"rentalmanagement/infrastructure/database/mappers"
	"rentalmanagement/logic/model"
)

// CustomerRepository implements the CustomerRepositoryInterface
type CustomerRepository struct {
	DB *gorm.DB
}

// NewCustomerRepository returns a new CustomerRepository instance
func NewCustomerRepository(db *gorm.DB) model.CustomerRepositoryInterface {
	return &CustomerRepository{
		DB: db,
	}
}

// AddCustomer adds a new customer to the database
func (repo *CustomerRepository) AddCustomer(customer model.Customer) (model.Customer, error) {
	var msg string
	customer.CustomerId = uuid.New().String()
	// Assuming you have a mappers package/function to convert a model to a persistence entity.
	customerPers := mappers.ConvertCustomerToCustomerPersistenceEntity(customer)

	// Attempt to add the customer to the database
	if err := repo.DB.Create(&customerPers).Error; err != nil {
		msg = fmt.Sprintf("Database Failed to add customer with Name %s", customer.Name)
		log.Error(msg, ": ", err)
		return customer, fmt.Errorf("%s: %w", msg, err)
	}
	// Convert the persistence entity back to the model if needed (e.g., to get the ID assigned by the DB)
	customer = mappers.ConvertCustomerPersistenceEntityToModel(customerPers)

	return customer, nil
}

// DeleteCustomer deletes a customer from the database using the customer's ID
func (repo *CustomerRepository) DeleteCustomer(customer model.Customer) error {
	var msg string
	// Delete the customer with the specified ID
	persistenceCustomer := mappers.ConvertCustomerToCustomerPersistenceEntity(customer)
	if tx := repo.DB.Delete(&entities.CustomerPersistenceEntity{}, persistenceCustomer.CustomerId); tx.Error != nil {
		//if err := repo.DB.Exec("DELETE FROM customer WHERE customer_id = ?", customerId); err != nil {
		msg = fmt.Sprintf("Database failed to delete customer with ID %s", customer.CustomerId)

		return fmt.Errorf("%s: %w", msg, tx.Error)
	}

	return nil
}

func (repo *CustomerRepository) GetCustomer(customerId string) (model.Customer, error) {
	var customerPers entities.CustomerPersistenceEntity
	var msg string

	// Try to find a customer with the given ID
	err := repo.DB.Where("customer_id = ?", customerId).First(&customerPers).Error
	if err != nil {
		// If the error is a record not found error, then return false without an error
		if err == gorm.ErrRecordNotFound {
			return model.Customer{}, nil
		}

		// For other database errors, log and return them
		msg = fmt.Sprintf("Database error when trying to fetch customer with ID %s", customerId)
		log.Error(msg, ": ", err)
		return model.Customer{}, fmt.Errorf("%s: %w", msg, err)
	}

	// If no errors occurred, it means a customer with the given ID was found
	return mappers.ConvertCustomerPersistenceEntityToModel(customerPers), nil
}
