package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	userpb "BEKS777/ADV2/proto/user"
)

type userService struct {
	userpb.UnimplementedUserServiceServer
}

func (s *userService) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	// Business logic to register user
	// ...
	// Generate user ID
	userID := "user123"

	res := &userpb.RegisterUserResponse{
		UserId: userID,
	}

	return res, nil
}

func (s *userService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	// Business logic to get user by ID
	// ...
	name := "Bekslan Yermekbayev"
	email := "assamyellow@gmail.com"

	res := &userpb.GetUserResponse{
		Name:  name,
		Email: email,
	}

	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, &userService{})

	fmt.Println("User service is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
