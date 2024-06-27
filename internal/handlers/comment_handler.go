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

// HandleComments yorum isteklerini yönetir
func (h *CommentHandler) HandleComments(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // switch-case ile istek metoduna ağırlık verir
	case http.MethodPost: // POST istegise yorum oluşturur
		h.createComment(w, r)
	case http.MethodGet: // GET istegise yorumları alır
		id := r.URL.Query().Get("id") // id parametresini alır
		if id != "" {                 // id parametresi varsa
			h.getCommentByID(w, r, id) // id parametresi ile yorumu alır
		} else { // id parametresi yoksa
			h.getCommentsByPostID(w, r) // post_id parametresi ile yorumları alır
		}
	case http.MethodPut: // PUT istegise yorumu günceller
		h.updateComment(w, r) // id parametresi ile yorumu günceller
	case http.MethodDelete: // DELETE istegise yorumu siler
		h.deleteComment(w, r) // id parametresi ile yorumu siler
	default: // istek metoduna ağırlık verilmez
		w.WriteHeader(http.StatusMethodNotAllowed) // istek metoduna ağırlık verir
	}
}

// createComment yorum oluşturur
func (h *CommentHandler) createComment(w http.ResponseWriter, r *http.Request) {
	var comment entities.Comment                    // yorum nesnesi oluşturur
	err := json.NewDecoder(r.Body).Decode(&comment) // json ile yorumu alır
	if err != nil {                                 // json ile yorumu alamazsa
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.CreateComment(&comment) // yorumu veritabanına kaydeder
	if err != nil {                         // veritabanına kaydederken hata varsa
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusCreated) // istek metoduna ağırlık verir
}

// getCommentByID yorumu alır
func (h *CommentHandler) getCommentByID(w http.ResponseWriter, r *http.Request, id string) {
	commentID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {                  // id parametresi alamazsa
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	comment, err := h.service.GetCommentByID(commentID) // id parametresi ile yorumu alır
	if err != nil {                                     // id parametresi ile yorumu alamazsa
		w.WriteHeader(http.StatusNotFound) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(comment) // yorumu json ile döndürür
}

// getCommentsByPostID yorumları alır
func (h *CommentHandler) getCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("post_id") // post_id parametresini alır
	if postID == "" {                      // post_id parametresi yoksa
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	postUUID, err := uuid.Parse(postID) // post_id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	comments, err := h.service.GetCommentsByPostID(postUUID) // post_id parametresi ile yorumları alır
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(comments)
}

// updateComment yorumu günceller
func (h *CommentHandler) updateComment(w http.ResponseWriter, r *http.Request) {
	var comment entities.Comment                    // yorum nesnesi oluşturur
	err := json.NewDecoder(r.Body).Decode(&comment) // json ile yorumu alır
	if err != nil {                                 // json ile yorumu alamazsa
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.UpdateComment(&comment) // yorumu veritabanına kaydeder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusOK) // istek metoduna ağırlık verir
}

// deleteComment yorumu siler
func (h *CommentHandler) deleteComment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // id parametresini alır
	if id == "" {                 // id parametresi yoksa
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	commentID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.DeleteComment(commentID) // id parametresi ile yorumu siler
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusOK) // istek metoduna ağırlık verir
}
