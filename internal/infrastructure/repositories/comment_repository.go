package repositories

import (
	"database/sql"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/google/uuid"
	"log"
)

// CommentRepository arayüz yorum işlemleri için yöntemleri tanımlar
type CommentRepository interface {
	Create(comment *entities.Comment) error
	GetByID(id uuid.UUID) (*entities.Comment, error)
	GetByPostID(postID uuid.UUID) ([]*entities.Comment, error)
	Update(comment *entities.Comment) error
	Delete(id uuid.UUID) error
}

// commentRepository yapı CommentRepository arayüzünü uygular
type commentRepository struct {
	db *sql.DB
}

// NewCommentRepository yeni bir CommentRepository oluşturur
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db}
}

// Create veritabanında yeni bir yorum oluşturur
func (r *commentRepository) Create(comment *entities.Comment) error {
	query := `INSERT INTO comments (id, post_id, user_id, content, created_at) VALUES ($1, $2, $3, $4, $5)`    // SQL sorgusu
	_, err := r.db.Exec(query, comment.ID, comment.PostID, comment.UserID, comment.Content, comment.CreatedAt) // SQL sorgusunu veritabanına yollar
	if err != nil {                                                                                            // hata varsa
		log.Printf("Error creating comment: %v", err)
		return err
	}
	return nil
}

// GetByID veritabanından kimliğe göre bir yorumu alır
func (r *commentRepository) GetByID(id uuid.UUID) (*entities.Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at FROM comments WHERE id = $1`
	row := r.db.QueryRow(query, id)
	comment := &entities.Comment{}
	err := row.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No comment found
		}
		log.Printf("Error getting comment by ID: %v", err)
		return nil, err
	}
	return comment, nil
}

// GetByPostID bir gönderiye ilişkin tüm yorumları veritabanından alır
func (r *commentRepository) GetByPostID(postID uuid.UUID) ([]*entities.Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at FROM comments WHERE post_id = $1`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		log.Printf("Error getting comments by post ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	comments := []*entities.Comment{}
	for rows.Next() {
		comment := &entities.Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt)
		if err != nil {
			log.Printf("Error scanning comment row: %v", err)
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// Update veritabanındaki bir yorumu günceller
func (r *commentRepository) Update(comment *entities.Comment) error {
	query := `UPDATE comments SET content = $1 WHERE id = $2`
	_, err := r.db.Exec(query, comment.Content, comment.ID)
	if err != nil {
		log.Printf("Error updating comment: %v", err)
		return err
	}
	return nil
}

// Delete veritabanından bir yorumu siler
func (r *commentRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting comment: %v", err)
		return err
	}
	return nil
}
