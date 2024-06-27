package services

import (
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

type CommentService interface {
	CreateComment(comment *entities.Comment) error
	GetCommentByID(id uuid.UUID) (*entities.Comment, error)
	GetCommentsByPostID(postID uuid.UUID) ([]*entities.Comment, error)
	UpdateComment(comment *entities.Comment) error
	DeleteComment(id uuid.UUID) error
}

// commentService arayüz yorum işlemleri için yöntemleri tanımlar
type commentService struct {
	repo repositories.CommentRepository
}

// NewCommentService yeni bir CommentService oluşturur
func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo}
}

// CreateComment veritabanında yeni bir yorum oluşturur
func (s *commentService) CreateComment(comment *entities.Comment) error {
	// Implement the create comment logic
	return s.repo.Create(comment)
}

// GetCommentByID veritabanından kimliğe göre bir yorumu alır
func (s *commentService) GetCommentByID(id uuid.UUID) (*entities.Comment, error) {
	// Implement the get comment by ID logic
	return s.repo.GetByID(id)
}

// GetCommentsByPostID bir gönderiye ilişkin tüm yorumları veritabanından alır
func (s *commentService) GetCommentsByPostID(postID uuid.UUID) ([]*entities.Comment, error) {
	// Implement the get comments by post ID logic
	return s.repo.GetByPostID(postID)
}

// UpdateComment veritabanında kimliğe göre bir yorumu günceller
func (s *commentService) UpdateComment(comment *entities.Comment) error {
	// Implement the update comment logic
	return s.repo.Update(comment)
}

// DeleteComment veritabanından kimliğe göre bir yorumu siler
func (s *commentService) DeleteComment(id uuid.UUID) error {
	// Implement the delete comment logic
	return s.repo.Delete(id)
}
