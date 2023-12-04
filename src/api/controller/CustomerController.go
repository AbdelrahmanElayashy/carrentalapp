package controller

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"rentalmanagement/api/controller/mappers"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/logic/operations"
)

type CustomerController struct {
	ops operations.CustomerOperationsInterface
	pb.UnimplementedCustomerServiceServer
}

func NewCustomerController(ops operations.CustomerOperationsInterface) CustomerController {
	return CustomerController{ops: ops}
}

// Implement the RentCar RPC method
func (controller CustomerController) RentCar(ctx context.Context, req *pb.RentCarRequest) (*pb.RentCarResponse, error) {
	log.Info("Starting to add a new rental.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.RentCarResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.Vin.GetVin() == "" || req.GetCustomerId() == "" || req.GetStartDate() == nil || req.GetEndDate() == nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("VIN, start date, customerId or end date is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.RentCarResponse{
			Error: errorDetail,
		}, nil
	}

	rental, err := controller.ops.RentCar(mappers.ConvertProtobufTimeStampToDate(req.StartDate), mappers.ConvertProtobufTimeStampToDate(req.EndDate), req.Vin.Vin, req.CustomerId)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.RentCarResponse{
			Error: errorDetail,
		}, nil
	}

	return &pb.RentCarResponse{
		Rental: mappers.ConvertModelRentalToProtobufRental(rental),
	}, nil
}

// Implement the CancelRental RPC method
func (controller CustomerController) CancelRental(ctx context.Context, req *pb.CancelRentalRequest) (*pb.CancelRentalResponse, error) {
	log.Info("Starting to cancel a rental.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.CancelRentalResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			},
			Result: false}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.RentalId == "" {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("RentalId is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.CancelRentalResponse{
			Error:  errorDetail,
			Result: false,
		}, nil
	}

	// Call the relevant logic operation to cancel the rental.
	err := controller.ops.CancelRental(req.RentalId)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.CancelRentalResponse{
			Error:  errorDetail,
			Result: false,
		}, nil
	}

	// If successful, return a response confirming the cancellation.
	return &pb.CancelRentalResponse{Result: true}, nil
}

// Implement the ListRentals RPC method
func (controller CustomerController) ListRentals(ctx context.Context, req *pb.ListRentalsRequest) (*pb.ListRentalsResponse, error) {
	log.Info("Starting to list rentals for a customer.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.ListRentalsResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.CustomerId == "" {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: "CustomerId is not valid",
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.ListRentalsResponse{
			Error: errorDetail,
		}, nil
	}

	// Invoke the relevant operation to list all rentals for the given customer ID.
	rentals, err := controller.ops.ListRentals(req.CustomerId)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.ListRentalsResponse{
			Error: errorDetail,
		}, nil
	}

	// Map the list of rentals to the protobuf response.
	var pbRentals []*pb.Rental
	for _, rental := range rentals {
		pbRental := mappers.ConvertModelRentalToProtobufRental(rental)
		pbRentals = append(pbRentals, pbRental)
	}
	return &pb.ListRentalsResponse{
		Rentals: pbRentals,
	}, nil
}
