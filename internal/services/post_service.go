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

type postService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) CreatePost(post *entities.Post) error {
	// Implement the create post logic
	return s.repo.Create(post)
}

func (s *postService) GetPostByID(id uuid.UUID) (*entities.Post, error) {
	// Implement the get post by ID logic
	return s.repo.GetByID(id)
}

func (s *postService) GetAllPosts() ([]*entities.Post, error) {
	// Implement the get all posts logic
	return s.repo.GetAll()
}

func (s *postService) UpdatePost(post *entities.Post) error {
	// Implement the update post logic
	return s.repo.Update(post)
}

func (s *postService) DeletePost(id uuid.UUID) error {
	// Implement the delete post logic
	return s.repo.Delete(id)
}
