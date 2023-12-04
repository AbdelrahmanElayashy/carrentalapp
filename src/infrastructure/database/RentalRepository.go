package database

import (
	"fmt"
	"github.com/google/uuid"
	"rentalmanagement/infrastructure/database/entities"
	"rentalmanagement/infrastructure/database/mappers"
	"rentalmanagement/logic/model"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RentalRepository implements the RentalRepositoryInterface
type RentalRepository struct {
	DB *gorm.DB
}

// NewRentalRepository returns a new RentalRepository instance
func NewRentalRepository(db *gorm.DB) model.RentalRepositoryInterface {
	return &RentalRepository{
		DB: db,
	}
}

// AddRental adds a new rental to the database
func (repo *RentalRepository) AddRental(rental model.Rental) (model.Rental, error) {
	var msg string
	rental.RentalId = uuid.New().String()
	rentalPres := mappers.ConvertRentalToRentalPersistenceEntity(rental)
	if err := repo.DB.Create(&rentalPres).Error; err != nil {
		msg = fmt.Sprintf("Database Failed to add rental for car with VIN %s", rental.Vin.Vin)
		log.Error(msg, ": ", err)
		return rental, fmt.Errorf("%s: %w", msg, err)
	}
	return mappers.ConvertRentalPersistenceEntityToRental(rentalPres), nil
}

// IsCarAvailableForRental checks if a car with the given VIN is available for rental between startDate and endDate
func (repo *RentalRepository) IsCarAvailableForRental(vin string, startDate, endDate time.Time) (bool, error) {
	var msg string
	var count int64
	if err := repo.DB.Model(&entities.RentalPersistenceEntity{}).
		Where("vin = ?", vin).
		Where("(start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?)", startDate, endDate, startDate, endDate).
		Count(&count).Error; err != nil {
		msg = fmt.Sprintf("Database Failed to check availability for car with VIN %s", vin)
		log.Error(msg, ": ", err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}

	if count > 0 {
		log.Warn(fmt.Sprintf("Car with VIN %s is not available for the given date range", vin))
		return false, nil
	}
	return true, nil
}

func (repo *RentalRepository) ListRentalsByCustomerId(customerId string) ([]model.Rental, error) {
	var rentalPersEntities []entities.RentalPersistenceEntity
	var rentals []model.Rental

	// Query for rentals by customerId
	if err := repo.DB.Where("customer_id = ?", customerId).Find(&rentalPersEntities).Error; err != nil {
		msg := fmt.Sprintf("Database failed to list rentals for customer with ID %s", customerId)
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	// Convert persistence entities to domain entities
	for _, rentalPers := range rentalPersEntities {
		rental := mappers.ConvertRentalPersistenceEntityToRental(rentalPers) // Assuming you have a function like this
		rentals = append(rentals, rental)
	}

	return rentals, nil
}

func (repo *RentalRepository) DeleteRental(rentalId string) error {
	var msg string

	// Directly delete the rental based on the ID
	if err := repo.DB.Where("rental_id = ?", rentalId).Delete(&entities.RentalPersistenceEntity{}).Error; err != nil {
		msg = fmt.Sprintf("Database failed to delete rental with ID %s", rentalId)
		log.Error(msg, ": ", err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}

func (repo *RentalRepository) DoesRentalExist(rentalId string) (bool, error) {
	var msg string
	var rentalPers entities.RentalPersistenceEntity

	// Check if the rental exists
	if err := repo.DB.Where("rental_id = ?", rentalId).First(&rentalPers).Error; err != nil {
		msg = fmt.Sprintf("Database failed to find rental with ID %s", rentalId)
		log.Error(msg, ": ", err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}

	// If we get to this point, it means a rental with the given ID exists
	return true, nil
}
