package main

import (
	"context"
	"fmt"
	"log"
	"net"

	apartmentpb "BEKS777/ADV2/proto/apartment"
)

type apartmentService struct {
	apartmentpb.UnimplementedApartmentServiceServer
}

func (s *apartmentService) GetAvailableApartments(ctx context.Context, req *apartmentpb.GetAvailableApartmentsRequest) (*apartmentpb.GetAvailableApartmentsResponse, error) {
	// Business logic to get available apartments in a city
	// ...
	apartments := []*apartmentpb.Apartment{
		{
			Id:    "home123",
			City: "Aktau",
			Room:  "two",
		},
		{
			Id:    "car456",
			City: "Astana"
			Room: "three"
			
		},
	}

	res := &apartmentpb.GetAvailableApartmentsResponse{
		Apartments: apartments,
	}

	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	apartmentpb.RegisterApartmentServiceServer(server, &apartmentService{})

	fmt.Println("Apartment service is running on port 50052...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
