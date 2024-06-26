package handlers

import (
	"encoding/json"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/services"
	"github.com/google/uuid"
	"net/http"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentHandler(service services.CommentService) *CommentHandler {
	return &CommentHandler{service}
}

func (h *CommentHandler) HandleComments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createComment(w, r)
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id != "" {
			h.getCommentByID(w, r, id)
		} else {
			h.getCommentsByPostID(w, r)
		}
	case http.MethodPut:
		h.updateComment(w, r)
	case http.MethodDelete:
		h.deleteComment(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *CommentHandler) createComment(w http.ResponseWriter, r *http.Request) {
	var comment entities.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.CreateComment(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CommentHandler) getCommentByID(w http.ResponseWriter, r *http.Request, id string) {
	commentID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment, err := h.service.GetCommentByID(commentID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(comment)
}

func (h *CommentHandler) getCommentsByPostID(w http.ResponseWriter, r *http.Request) {
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
	comments, err := h.service.GetCommentsByPostID(postUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comments)
}

func (h *CommentHandler) updateComment(w http.ResponseWriter, r *http.Request) {
	var comment entities.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.UpdateComment(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CommentHandler) deleteComment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	commentID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.DeleteComment(commentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
