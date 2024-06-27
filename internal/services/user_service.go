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

// NewUserService yeni bir UserService oluşturur
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

// CreateUser veritabanında yeni bir kullanıcı oluşturur
func (s *userService) CreateUser(user *entities.User) error {
	return s.repo.Create(user)
}

// GetUserByID veritabanından kimliğe göre bir kullanıcıyı alır
func (s *userService) GetUserByID(id uuid.UUID) (*entities.User, error) {
	return s.repo.GetByID(id)
}

// GetUserByUsername veritabanından kimliğe göre bir kullanıcıyı alır
func (s *userService) GetUserByUsername(username string) (*entities.User, error) {
	return s.repo.GetByUsername(username)
}

// UpdateUser veritabanında kimliğe göre bir kullanıcıyı günceller
func (s *userService) UpdateUser(user *entities.User) error {
	return s.repo.Update(user)
}

// DeleteUser veritabanından kimliğe göre bir kullanıcıyı siler
func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}
