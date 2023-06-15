package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	
	userpb "BEKS777/ADV2/proto/user"
)

func main() {
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	userClient := userpb.NewUserServiceClient(userConn)


	// User service - RegisterUser
	registerReq := &userpb.RegisterUserRequest{
		Name:  "Bekslan Yermekbayev",
		Email: "assamyellow@gmail.com",
	}

	registerRes, err := userClient.RegisterUser(context.Background(), registerReq)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	fmt.Println("User ID:", registerRes.UserId)
}
