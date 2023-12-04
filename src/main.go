package main

import (
	"fmt"
	"net"
	"os"
	"rentalmanagement/api/controller"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/infrastructure/database"
	"rentalmanagement/infrastructure/external"
	"rentalmanagement/logic/operations"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Create SQLite in-memory database
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	// Migrate models to database tables
	err = database.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
		return
	}

	// Start a gRPC server
	portEnv := os.Getenv("GRPC_PORT")
	if portEnv == "" {
		portEnv = "80"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", portEnv))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	carAPIEnv := os.Getenv("CAR_API_URL")
	if carAPIEnv == "" {
		carAPIEnv = "http://dm-car:80"
	}

	rentalRepo := database.NewRentalRepository(database.InMemDB)
	customerRepo := database.NewCustomerRepository(database.InMemDB)
	carApi := external.NewCarAPI(carAPIEnv)

	rentalOps := operations.NewRentalsCollectionOperations(rentalRepo, carApi)
	customerOps := operations.NewCustomerOperations(rentalRepo, carApi)
	customerCollectionOps := operations.NewCustomersCollectionOperations(customerRepo)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterRentalsCollectionServiceServer(grpcServer, controller.NewRentalsCollectionController(rentalOps))
	pb.RegisterCustomerServiceServer(grpcServer, controller.NewCustomerController(customerOps))
	pb.RegisterCustomersCollectionServiceServer(grpcServer, controller.NewCustomersCollectionController(customerCollectionOps))

	log.Println(fmt.Sprintf("Server is running on port %s", portEnv))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
