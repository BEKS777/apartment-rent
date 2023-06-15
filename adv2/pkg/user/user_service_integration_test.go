package user

import (
	"fmt"
	"testing"
)

func TestUserServiceIntegration_RegisterAndGetUser(t *testing.T) {
	// Start the UserService server
	userService := NewUserService()
	go userService.StartServer()

	// Register a user
	userID, err := userService.RegisterUser("Bekslan Yermekbayev", "assamyellow@gmail.com")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Get the registered user
	user, err := userService.GetUser(userID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	// Verify the user details
	if user.Name != "Bekslan Yermekbayev" || user.Email != "assamyellow@gmail.com" {
		t.Errorf("Unexpected user details. Expected: 'Bekslan Yermekbayev', 'assamyellow@gmail.com', Got: '%s', '%s'", user.Name, user.Email)
	}

	// Stop the UserService server
	userService.StopServer()
}

func TestUserServiceIntegration_RegisterAndGetUser_NotFound(t *testing.T) {
	// Start the UserService server
	userService := NewUserService()
	go userService.StartServer()

	// Get a non-existent user
	_, err := userService.GetUser("nonexistent")
	if err == nil {
		t.Error("Expected error when getting non-existent user, but got nil")
	}

	// Stop the UserService server
	userService.StopServer()
}

func TestUserServiceIntegration_GetAllUsers(t *testing.T) {
	// Start the UserService server
	userService := NewUserService()
	go userService.StartServer()

	// Register some users
	_, err := userService.RegisterUser("Bekslan Yermekbayev", "assamyellow@gmail.com")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}
	_, err = userService.RegisterUser("Otesh Didar", "didaro@gmail.com")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Get all users
	users := userService.GetAllUsers()

	// Verify the number of users
	expectedCount := 2
	if len(users) != expectedCount {
		t.Errorf("Unexpected number of users. Expected: %d, Got: %d", expectedCount, len(users))
	}

	// Print the users
	fmt.Println("All Users:")
	for _, user := range users {
		fmt.Printf("- Name: %s, Email: %s\n", user.Name, user.Email)
	}

	// Stop the UserService server
	userService.StopServer()
}
