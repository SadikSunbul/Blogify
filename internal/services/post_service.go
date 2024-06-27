package services

import (
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

type PostService interface {
	CreatePost(post *entities.Post) error
	GetPostByID(id uuid.UUID) (*entities.Post, error)
	GetAllPosts() ([]*entities.Post, error)
	UpdatePost(post *entities.Post) error
	DeletePost(id uuid.UUID) error
}

// postService arayüz yorum işlemleri için yöntemleri tanımlar
type postService struct {
	repo repositories.PostRepository
}

// NewPostService yeni bir PostService oluşturur
func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo}
}

// CreatePost veritabanında yeni bir yorum oluşturur
func (s *postService) CreatePost(post *entities.Post) error {
	// Implement the create post logic
	return s.repo.Create(post)
}

// GetPostByID veritabanından kimliğe göre bir yorumu alır
func (s *postService) GetPostByID(id uuid.UUID) (*entities.Post, error) {
	// Implement the get post by ID logic
	return s.repo.GetByID(id)
}

// GetAllPosts veritabanından tüm yorumları alır
func (s *postService) GetAllPosts() ([]*entities.Post, error) {
	// Implement the get all posts logic
	return s.repo.GetAll()
}

// UpdatePost veritabanında kimliğe göre bir yorumu günceller
func (s *postService) UpdatePost(post *entities.Post) error {
	// Implement the update post logic
	return s.repo.Update(post)
}

// DeletePost veritabanından kimliğe göre bir yorumu siler
func (s *postService) DeletePost(id uuid.UUID) error {
	// Implement the delete post logic
	return s.repo.Delete(id)
}
