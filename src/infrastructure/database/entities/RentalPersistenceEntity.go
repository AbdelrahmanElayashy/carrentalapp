package entities

import (
	"time"

	"github.com/google/uuid"
)

type RentalPersistenceEntity struct {
	RentalId   uuid.UUID `gorm:"type:uuid;primary_key;"`
	StartDate  time.Time
	EndDate    time.Time
	CustomerId string `gorm:"type:text;"`
	Vin        string `gorm:"type:text;"`
}

// TableName sets the insert table name for this struct type
func (p *RentalPersistenceEntity) TableName() string {
	return "rental"
}
