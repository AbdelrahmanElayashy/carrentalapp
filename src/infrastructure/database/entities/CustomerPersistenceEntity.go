package entities

import "github.com/google/uuid"

type CustomerPersistenceEntity struct {
	CustomerId uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name       string    `gorm:"text"`
}

// TableName sets the insert table name for this struct type
func (p *CustomerPersistenceEntity) TableName() string {
	return "customer"
}
