package repositories

import (
	"database/sql"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/google/uuid"
	"log"
)

// PostRepository arayüz post operasyonlar için yöntemleri tanımlar
type PostRepository interface {
	Create(post *entities.Post) error
	GetByID(id uuid.UUID) (*entities.Post, error)
	GetAll() ([]*entities.Post, error)
	Update(post *entities.Post) error
	Delete(id uuid.UUID) error
}

// postRepository yapı PostRepository arayüzünü uygular
type postRepository struct {
	db *sql.DB
}

// NewPostRepository yeni bir PostRepository oluşturur
func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db}
}

// Create veritabanında yeni bir gönderi oluşturur
func (r *postRepository) Create(post *entities.Post) error {
	query := `INSERT INTO posts (id, user_id, title, content, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, post.ID, post.UserID, post.Title, post.Content, post.CreatedAt)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		return err
	}
	return nil
}

// GetByID veritabanından kimliğe göre bir gönderiyi alır
func (r *postRepository) GetByID(id uuid.UUID) (*entities.Post, error) {
	query := `SELECT id, user_id, title, content, created_at FROM posts WHERE id = $1`
	row := r.db.QueryRow(query, id)
	post := &entities.Post{}
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No post found
		}
		log.Printf("Error getting post by ID: %v", err)
		return nil, err
	}
	return post, nil
}

// GetAll veritabanındaki tüm gönderileri alır
func (r *postRepository) GetAll() ([]*entities.Post, error) {
	query := `SELECT id, user_id, title, content, created_at FROM posts`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error getting all posts: %v", err)
		return nil, err
	}
	defer rows.Close()

	posts := []*entities.Post{}
	for rows.Next() {
		post := &entities.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			log.Printf("Error scanning post row: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Update veritabanındaki bir gönderiyi günceller
func (r *postRepository) Update(post *entities.Post) error {
	query := `UPDATE posts SET title = $1, content = $2 WHERE id = $3`
	_, err := r.db.Exec(query, post.Title, post.Content, post.ID)
	if err != nil {
		log.Printf("Error updating post: %v", err)
		return err
	}
	return nil
}

// Delete veri tabanından bir gönderiyi siler
func (r *postRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting post: %v", err)
		return err
	}
	return nil
}
