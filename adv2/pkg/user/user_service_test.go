package user

import (
	"testing"
)

func TestUserService_RegisterUser(t *testing.T) {
	// Create a new instance of UserService
	userService := NewUserService()

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
}

func TestUserService_GetUser(t *testing.T) {
	// Create a new instance of UserService
	userService := NewUserService()

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

	// Verify the user ID
	if user.ID != userID {
		t.Errorf("Unexpected user ID. Expected: '%s', Got: '%s'", userID, user.ID)
	}
}

func TestUserService_GetUser_NotFound(t *testing.T) {
	// Create a new instance of UserService
	userService := NewUserService()

	// Get a non-existent user
	_, err := userService.GetUser("nonexistent")
	if err == nil {
		t.Error("Expected error when getting non-existent user, but got nil")
	}
}

func TestUserService_GetUser_AllUsers(t *testing.T) {
	// Create a new instance of UserService
	userService := NewUserService()

	// Register some users
	_, err := userService.RegisterUser("Yermekbayev Bekslan", "assamyellow@gmail.com")
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
}
