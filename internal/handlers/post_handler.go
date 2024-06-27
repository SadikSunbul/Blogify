package handlers

import (
	"encoding/json"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/services"
	"github.com/google/uuid"
	"net/http"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{service}
}

// HandlePosts post isteklerini yönetir
func (h *PostHandler) HandlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // switch-case ile istek metoduna ağırlık verir
	case http.MethodPost: // POST istegise post oluşturur
		h.createPost(w, r) // postu veritabanına kaydeder
	case http.MethodGet: // GET istegise postları alır
		id := r.URL.Query().Get("id") // id parametresini alır
		if id != "" {                 // id parametresi varsa
			h.getPostByID(w, r, id) // id parametresi ile postu alır
		} else { // id parametresi yoksa
			h.getAllPosts(w, r) // postları alır
		}
	case http.MethodPut: // PUT istegise postu günceller
		h.updatePost(w, r) // id parametresi ile postu günceller
	case http.MethodDelete: // DELETE istegise postu siler
		h.deletePost(w, r) // id parametresi ile postu siler
	default: // istek metoduna ağırlık verilmez
		w.WriteHeader(http.StatusMethodNotAllowed) // istek metoduna ağırlık verir
	}
}

// createPost post oluşturur
func (h *PostHandler) createPost(w http.ResponseWriter, r *http.Request) {
	var post entities.Post                       // post nesnesi oluşturur
	err := json.NewDecoder(r.Body).Decode(&post) // json ile postu alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.CreatePost(&post) // postu veritabanına kaydeder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusCreated) // istek metoduna ağırlık verir
}

// getPostByID postu alır
func (h *PostHandler) getPostByID(w http.ResponseWriter, r *http.Request, id string) {
	postID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	post, err := h.service.GetPostByID(postID) // id parametresi ile postu alır
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(post) // postu json ile döndürür
}

// getAllPosts postları alır
func (h *PostHandler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAllPosts() // postları alır
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(posts) // postları json ile döndürür
}

// updatePost postu günceller
func (h *PostHandler) updatePost(w http.ResponseWriter, r *http.Request) {
	var post entities.Post                       // post nesnesi oluşturur
	err := json.NewDecoder(r.Body).Decode(&post) // json ile postu alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.UpdatePost(&post) // postu veritabanına kaydeder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusOK) // istek metoduna ağırlık verir
}

// deletePost postu siler
func (h *PostHandler) deletePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // id parametresini alır
	if id == "" {                 // id parametresi yoksa
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	postID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.DeletePost(postID) // id parametresi ile postu siler
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusOK) // istek metoduna ağırlık verir
}
