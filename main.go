package main

import (
	"fmt"
	"log"

	"github.com/BEKS777/apartment-rent/pkg/user"
)

func main() {
	userService := user.NewUserService()

	// Register a user
	userID, err := userService.RegisterUser("Bekslan Yermekbayev", "assamyellow@gmail.com")
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	fmt.Println("User ID:", userID)

	// Get user details
	userDetails, err := userService.GetUser(userID)
	if err != nil {
		log.Fatalf("Failed to get user details: %v", err)
	}

	fmt.Printf("User Details: %+v\n", userDetails)
}
