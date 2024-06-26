package repositories

import (
	"database/sql"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/google/uuid"
	"log"
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
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

// Create veritabanında yeni bir kullanıcı oluşturur
func (r *userRepository) Create(user *entities.User) error {
	query := `INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

// GetByID veritabanından kullanıcıyı kimliğine göre alır
func (r *userRepository) GetByID(id uuid.UUID) (*entities.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}
	return user, nil
}

// GetByUsername retrieves a user by username from the database
func (r *userRepository) GetByUsername(username string) (*entities.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE username = $1`
	row := r.db.QueryRow(query, username)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		log.Printf("Error getting user by username: %v", err)
		return nil, err
	}
	return user, nil
}

// Update updates a user in the database
func (r *userRepository) Update(user *entities.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

// Delete deletes a user from the database
func (r *userRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
