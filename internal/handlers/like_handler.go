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

// HandleLikes yorum isteklerini yönetir
func (h *LikeHandler) HandleLikes(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // switch-case ile istek metoduna ağırlık verir
	case http.MethodPost: // POST istegise yorum oluşturur
		h.createLike(w, r) // yorumu veritabanına kaydeder
	case http.MethodGet: // GET istegise yorumları alır
		id := r.URL.Query().Get("id") // id parametresini alır
		if id != "" {                 // id parametresi varsa
			h.getLikeByID(w, r, id) // id parametresi ile yorumu alır
		} else { // id parametresi yoksa
			h.getLikesByPostID(w, r) // post_id parametresi ile yorumları alır
		}
	case http.MethodDelete: // DELETE istegise yorumu siler
		h.deleteLike(w, r) // id parametresi ile yorumu siler
	default: // istek metoduna ağırlık verilmez
		w.WriteHeader(http.StatusMethodNotAllowed) // istek metoduna ağırlık verir
	}
}

// createLike yorum oluşturur
func (h *LikeHandler) createLike(w http.ResponseWriter, r *http.Request) {
	var like entities.Like                       // yorum nesnesi oluşturur
	err := json.NewDecoder(r.Body).Decode(&like) // json ile yorumu alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.CreateLike(&like) // yorumu veritabanına kaydeder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusCreated) // istek metoduna ağırlık verir
}

// getLikeByID yorumu alır
func (h *LikeHandler) getLikeByID(w http.ResponseWriter, r *http.Request, id string) {
	likeID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	like, err := h.service.GetLikeByID(likeID) // id parametresi ile yorumu alır
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(like) // yorumu jsona dönüştürür
}

// getLikesByPostID yorumları alır
func (h *LikeHandler) getLikesByPostID(w http.ResponseWriter, r *http.Request) {
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
	likes, err := h.service.GetLikesByPostID(postUUID) // post_id parametresi ile yorumları alır
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(likes) // yorumları jsona dönüştürür
}

// deleteLike yorumu siler
func (h *LikeHandler) deleteLike(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // id parametresini alır
	if id == "" {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	likeID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.DeleteLike(likeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusOK) // istek metoduna ağırlık verir
}
