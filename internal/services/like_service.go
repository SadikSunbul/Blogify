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

// likeService arayüz benzer işlemler için yöntemleri tanımlar
type likeService struct {
	repo repositories.LikeRepository
}

// NewLikeService yeni bir LikeService oluşturur
func NewLikeService(repo repositories.LikeRepository) LikeService {
	return &likeService{repo}
}

// CreateLike veritabanında yeni bir beğeni oluşturur
func (s *likeService) CreateLike(like *entities.Like) error {
	// Implement the create like logic
	return s.repo.Create(like)
}

// GetLikeByID veritabanından kimliğe göre bir beğeni alır
func (s *likeService) GetLikeByID(id uuid.UUID) (*entities.Like, error) {
	return s.repo.GetByID(id)
}

// GetLikesByPostID bir gönderiye ilişkin tüm beğenileri veritabanından alır
func (s *likeService) GetLikesByPostID(postID uuid.UUID) ([]*entities.Like, error) {
	return s.repo.GetByPostID(postID)
}

// DeleteLike veritabanından kimliğe göre bir beğeni siler
func (s *likeService) DeleteLike(id uuid.UUID) error {
	return s.repo.Delete(id)
}
