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

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo}
}

func (s *commentService) CreateComment(comment *entities.Comment) error {
	// Implement the create comment logic
	return s.repo.Create(comment)
}

func (s *commentService) GetCommentByID(id uuid.UUID) (*entities.Comment, error) {
	// Implement the get comment by ID logic
	return s.repo.GetByID(id)
}

func (s *commentService) GetCommentsByPostID(postID uuid.UUID) ([]*entities.Comment, error) {
	// Implement the get comments by post ID logic
	return s.repo.GetByPostID(postID)
}

func (s *commentService) UpdateComment(comment *entities.Comment) error {
	// Implement the update comment logic
	return s.repo.Update(comment)
}

func (s *commentService) DeleteComment(id uuid.UUID) error {
	// Implement the delete comment logic
	return s.repo.Delete(id)
}
