package controller

import (
	"context"
	"fmt"
	"rentalmanagement/api/controller/mappers"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/logic/operations"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/codes"
)

type RentalsCollectionController struct {
	ops operations.RentalsCollectionOperationsInterface
	pb.UnimplementedRentalsCollectionServiceServer
}

func NewRentalsCollectionController(ops operations.RentalsCollectionOperationsInterface) RentalsCollectionController {
	return RentalsCollectionController{ops: ops}
}

// Implement the ListAvailableCars RPC method
func (controller RentalsCollectionController) ListAvailableCars(ctx context.Context, req *pb.ListAvailableCarsRequest) (*pb.ListAvailableCarsResponse, error) {

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.ListAvailableCarsResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.StartDate == nil || req.EndDate == nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("Start date, or end date is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.ListAvailableCarsResponse{
			Error: errorDetail,
		}, nil
	}

	cars, err := controller.ops.ListAvailableCars(mappers.ConvertProtobufTimeStampToDate(req.StartDate), mappers.ConvertProtobufTimeStampToDate(req.EndDate))
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.ListAvailableCarsResponse{
			Error: errorDetail,
		}, nil
	}

	pbCars := []*pb.Car{}
	for _, car := range cars {
		pbCars = append(pbCars, mappers.ConvertModelCarToProtobufCar(car))
	}

	return &pb.ListAvailableCarsResponse{
		Cars: pbCars,
	}, nil
}
