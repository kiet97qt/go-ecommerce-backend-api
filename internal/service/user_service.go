package service

import (
	"context"
	"fmt"

	"go-ecommerce-backend-api/internal/models"

	"github.com/google/uuid"
)

// UserService exposes read operations for user resources.
type UserService interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

type userService struct {
	users map[string]models.User
}

// NewUserService returns a mock user service populated with sample data.
func NewUserService() UserService {
	return &userService{
		users: map[string]models.User{
			"1": {ID: uuid.New(), Username: "Alice Nguyen"},
			"2": {ID: uuid.New(), Username: "Bob Tran"},
		},
	}
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %s not found", id)
	}
	return &user, nil
}
