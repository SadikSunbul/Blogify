package handlers

import (
	"encoding/json"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/services"
	"github.com/google/uuid"
	"net/http"
)

type LikeHandler struct {
	service services.LikeService
}

func NewLikeHandler(service services.LikeService) *LikeHandler {
	return &LikeHandler{service}
}

func (h *LikeHandler) HandleLikes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createLike(w, r)
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id != "" {
			h.getLikeByID(w, r, id)
		} else {
			h.getLikesByPostID(w, r)
		}
	case http.MethodDelete:
		h.deleteLike(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *LikeHandler) createLike(w http.ResponseWriter, r *http.Request) {
	var like entities.Like
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.CreateLike(&like)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *LikeHandler) getLikeByID(w http.ResponseWriter, r *http.Request, id string) {
	likeID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	like, err := h.service.GetLikeByID(likeID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(like)
}

func (h *LikeHandler) getLikesByPostID(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("post_id")
	if postID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	likes, err := h.service.GetLikesByPostID(postUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(likes)
}

func (h *LikeHandler) deleteLike(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	likeID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.DeleteLike(likeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
