package repositories

import (
	"database/sql"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/google/uuid"
	"log"
)

// LikeRepository arayüz benzer işlemler için yöntemleri tanımlar
type LikeRepository interface {
	Create(like *entities.Like) error
	GetByID(id uuid.UUID) (*entities.Like, error)
	GetByPostID(postID uuid.UUID) ([]*entities.Like, error)
	Delete(id uuid.UUID) error
}

// likeRepository yapı LikeRepository arayüzünü uygular
type likeRepository struct {
	db *sql.DB
}

// NewLikeRepository yeni bir LikeRepository oluşturur
func NewLikeRepository(db *sql.DB) LikeRepository {
	return &likeRepository{db}
}

// Create veritabanında yeni bir beğeni oluşturur
func (r *likeRepository) Create(like *entities.Like) error {
	query := `INSERT INTO likes (id, post_id, user_id) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, like.ID, like.PostID, like.UserID)
	if err != nil {
		log.Printf("Error creating like: %v", err)
		return err
	}
	return nil
}

// GetByID veritabanından kimliğe göre beğeniyi alır
func (r *likeRepository) GetByID(id uuid.UUID) (*entities.Like, error) {
	query := `SELECT id, post_id, user_id FROM likes WHERE id = $1`
	row := r.db.QueryRow(query, id)
	like := &entities.Like{}
	err := row.Scan(&like.ID, &like.PostID, &like.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No like found
		}
		log.Printf("Error getting like by ID: %v", err)
		return nil, err
	}
	return like, nil
}

// GetByPostID veritabanındaki bir gönderiye ilişkin tüm beğenileri alır
func (r *likeRepository) GetByPostID(postID uuid.UUID) ([]*entities.Like, error) {
	query := `SELECT id, post_id, user_id FROM likes WHERE post_id = $1`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		log.Printf("Error getting likes by post ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	likes := []*entities.Like{}
	for rows.Next() {
		like := &entities.Like{}
		err := rows.Scan(&like.ID, &like.PostID, &like.UserID)
		if err != nil {
			log.Printf("Error scanning like row: %v", err)
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

// Delete bir beğeniyi veri tabanından siler
func (r *likeRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM likes WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting like: %v", err)
		return err
	}
	return nil
}
