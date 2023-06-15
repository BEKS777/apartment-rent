package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	rentalpb "BEKS777/ADV2/proto/rental"
)

type rentalService struct {
	rentalpb.UnimplementedRentalServiceServer
}

func (s *rentalService) RentApartment(ctx context.Context, req *rentalpb.RentApartmentRequest) (*rentalpb.RentApartmentResponse, error) {
	rentalID := "rental123"

	res := &rentalpb.RentApartmentResponse{
		RentalId: rentalID,
	}

	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	rentalpb.RegisterRentalServiceServer(server, &rentalService{})

	fmt.Println("Rental service is running on port 50053...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
