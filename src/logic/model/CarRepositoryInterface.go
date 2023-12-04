package model

type CarRepositoryInterface interface {
	ListAllCars() ([]Car, error)
	CarExists(vin Vin) (bool, error)
}
