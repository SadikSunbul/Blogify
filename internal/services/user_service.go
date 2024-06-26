package services

import (
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *entities.User) error
	GetUserByID(id uuid.UUID) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *entities.User) error {
	// Implement the create user logic
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id uuid.UUID) (*entities.User, error) {
	// Implement the get user by ID logic
	return s.repo.GetByID(id)
}

func (s *userService) GetUserByUsername(username string) (*entities.User, error) {
	// Implement the get user by username logic
	return s.repo.GetByUsername(username)
}

func (s *userService) UpdateUser(user *entities.User) error {
	// Implement the update user logic
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	// Implement the delete user logic
	return s.repo.Delete(id)
}
