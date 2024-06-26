package entities

import "github.com/google/uuid"

type Like struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	PostID uuid.UUID `json:"post_id"`
}
