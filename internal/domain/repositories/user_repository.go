package repositories

import (
	"database/sql"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/google/uuid"
)

// UserRepository arayüzü kullanıcı işlemleri için yöntemleri tanımlar
type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uuid.UUID) (*entities.User, error)
	GetByUsername(username string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository yeni bir UserRepository oluşturur
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

// Create creates a new user in the database
func (r *userRepository) Create(user *entities.User) error {
	// Implement the create operation
	return nil
}

// GetByID retrieves a user by ID from the database
func (r *userRepository) GetByID(id uuid.UUID) (*entities.User, error) {
	// Implement the get by ID operation
	return nil, nil
}

// GetByUsername retrieves a user by username from the database
func (r *userRepository) GetByUsername(username string) (*entities.User, error) {
	// Implement the get by username operation
	return nil, nil
}

// Update updates a user in the database
func (r *userRepository) Update(user *entities.User) error {
	// Implement the update operation
	return nil
}

// Delete deletes a user from the database
func (r *userRepository) Delete(id uuid.UUID) error {
	// Implement the delete operation
	return nil
}
