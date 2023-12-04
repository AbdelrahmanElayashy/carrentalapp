package mappers

import (
	"rentalmanagement/infrastructure/database/entities"
	"rentalmanagement/logic/model"

	"github.com/google/uuid"
)

func ConvertRentalToRentalPersistenceEntity(rental model.Rental) entities.RentalPersistenceEntity {
	rentalUUID, err := uuid.Parse(rental.RentalId)
	if err != nil {
		panic("Could not convert PersistenceEntity")
	}
	return entities.RentalPersistenceEntity{
		RentalId:   rentalUUID,
		StartDate:  rental.StartDate,
		EndDate:    rental.EndDate,
		CustomerId: rental.CustomerId,
		Vin:        rental.Vin.Vin,
	}
}

func ConvertRentalPersistenceEntityToRental(rentalPers entities.RentalPersistenceEntity) model.Rental {

	return model.Rental{
		RentalId:   rentalPers.RentalId.String(),
		StartDate:  rentalPers.StartDate,
		EndDate:    rentalPers.EndDate,
		CustomerId: rentalPers.CustomerId,
		Vin:        model.Vin{Vin: rentalPers.Vin},
	}
}
