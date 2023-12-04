package mappers

import (
	"github.com/google/uuid"
	"rentalmanagement/infrastructure/database/entities"
	"rentalmanagement/logic/model"
)

func ConvertCustomerToCustomerPersistenceEntity(customer model.Customer) entities.CustomerPersistenceEntity {
	customerUUID, err := uuid.Parse(customer.CustomerId)
	if err != nil {
		panic("Could not convert customerid")
	}
	return entities.CustomerPersistenceEntity{
		CustomerId: customerUUID,
		Name:       customer.Name,
	}
}

func ConvertCustomerPersistenceEntityToModel(customerPers entities.CustomerPersistenceEntity) model.Customer {

	return model.Customer{
		CustomerId: customerPers.CustomerId.String(),
		Name:       customerPers.Name,
	}
}
