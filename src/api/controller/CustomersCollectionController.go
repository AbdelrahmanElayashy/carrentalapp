package controller

import (
	"context"
	"rentalmanagement/api/controller/mappers"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/logic/operations"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type CustomersCollectionController struct {
	ops operations.CustomersCollectionOperationsInterface
	pb.UnimplementedCustomersCollectionServiceServer
}

func NewCustomersCollectionController(ops operations.CustomersCollectionOperationsInterface) CustomersCollectionController {
	return CustomersCollectionController{ops: ops}
}

// Implement the RegisterCustomer RPC method
func (controller CustomersCollectionController) RegisterCustomer(ctx context.Context, req *pb.RegisterCustomerRequest) (*pb.RegisterCustomerResponse, error) {
	log.Info("Starting to register a new customer.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.RegisterCustomerResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	if req.Name == "" {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: "Customer name is not valid.",
		}
		log.Errorf("Error: %s", errorDetail.Details)
		return &pb.RegisterCustomerResponse{
			Error: errorDetail,
		}, nil
	}

	customer, err := controller.ops.RegisterCustomer(req.Name)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.RegisterCustomerResponse{
			Error: errorDetail,
		}, nil
	}

	return &pb.RegisterCustomerResponse{
		Customer: mappers.ConvertModelCustomerToProtobufCustomer(customer),
	}, nil
}

// Implement the DeregisterCustomer RPC method
func (controller CustomersCollectionController) DeregisterCustomer(ctx context.Context, req *pb.DeregisterCustomerRequest) (*pb.DeregisterCustomerResponse, error) {
	log.Info("Starting to deregister a customer.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.DeregisterCustomerResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	if req.CustomerId == "" {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: "CustomerId is not valid.",
		}
		log.Errorf("Error: %s", errorDetail.Details)
		return &pb.DeregisterCustomerResponse{
			Error: errorDetail,
		}, nil
	}

	err := controller.ops.DeregisterCustomer(req.CustomerId)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.DeregisterCustomerResponse{
			Error: errorDetail,
		}, nil
	}

	return &pb.DeregisterCustomerResponse{}, nil
}
