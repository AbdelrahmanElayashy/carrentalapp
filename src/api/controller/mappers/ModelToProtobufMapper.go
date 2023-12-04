package mappers

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/logic/model"
	"time"
)

func ConvertModelCarToProtobufCar(car model.Car) *pb.Car {
	return &pb.Car{
		Vin:         &pb.Vin{Vin: car.Vin.Vin},
		Model:       car.Model,
		Brand:       car.Brand,
		PricePerDay: int32(car.PricePerDay)}
}

func ConvertModelCustomerToProtobufCustomer(customer model.Customer) *pb.Customer {
	return &pb.Customer{
		CustomerId: customer.CustomerId,
		Name:       customer.Name,
	}
}

func ConvertModelRentalToProtobufRental(rental model.Rental) *pb.Rental {
	return &pb.Rental{
		RentalId:   rental.RentalId,
		CustomerId: rental.CustomerId,
		Vin:        &pb.Vin{Vin: rental.Vin.Vin},
		StartDate:  mapModelDateToProtobufTimestamp(rental.StartDate),
		EndDate:    mapModelDateToProtobufTimestamp(rental.EndDate),
	}
}

func mapModelDateToProtobufTimestamp(date time.Time) *timestamppb.Timestamp {
	return timestamppb.New(date)
}
