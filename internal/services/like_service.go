package services

import (
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/infrastructure/repositories"
	"github.com/google/uuid"
)

type LikeService interface {
	CreateLike(like *entities.Like) error
	GetLikeByID(id uuid.UUID) (*entities.Like, error)
	GetLikesByPostID(postID uuid.UUID) ([]*entities.Like, error)
	DeleteLike(id uuid.UUID) error
}

type likeService struct {
	repo repositories.LikeRepository
}

func NewLikeService(repo repositories.LikeRepository) LikeService {
	return &likeService{repo}
}

func (s *likeService) CreateLike(like *entities.Like) error {
	// Implement the create like logic
	return s.repo.Create(like)
}

func (s *likeService) GetLikeByID(id uuid.UUID) (*entities.Like, error) {
	// Implement the get like by ID logic
	return s.repo.GetByID(id)
}

func (s *likeService) GetLikesByPostID(postID uuid.UUID) ([]*entities.Like, error) {
	// Implement the get likes by post ID logic
	return s.repo.GetByPostID(postID)
}

func (s *likeService) DeleteLike(id uuid.UUID) error {
	// Implement the delete like logic
	return s.repo.Delete(id)
}
