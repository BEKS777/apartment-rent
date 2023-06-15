package user

import (
	"errors"
	"fmt"
)

type UserService struct {
	users map[string]User
}

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]User),
	}
}

func (s *UserService) RegisterUser(name, email string) (string, error) {
	// Check if user already exists with the same email
	for _, user := range s.users {
		if user.Email == email {
			return "", errors.New("user with the same email already exists")
		}
	}

	// Generate a unique user ID (for simplicity, using a sequential number)
	userID := fmt.Sprintf("user%d", len(s.users)+1)

	// Create a new user
	user := User{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	// Store the user
	s.users[userID] = user

	return userID, nil
}

func (s *UserService) GetUser(userID string) (*User, error) {
	// Check if user exists
	user, ok := s.users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
