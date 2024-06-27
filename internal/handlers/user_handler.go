package handlers

import (
	"encoding/json"
	"github.com/SadikSunbul/Blogify/internal/domain/entities"
	"github.com/SadikSunbul/Blogify/internal/services"
	"github.com/google/uuid"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

// HandleUsers kullanıcı isteklerini yönetir
func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // switch-case ile istek metoduna ağırlık verir
	case http.MethodPost: // POST istegise kullanıcı oluşturur
		h.createUser(w, r) // kullanıcıyı veritabanına kaydeder
	case http.MethodGet: // GET istegise kullanıcıları alır
		id := r.URL.Query().Get("id") // id parametresini alır
		if id != "" {                 // id parametresi varsa
			h.getUserByID(w, r, id) // id parametresi ile kullanıcıyı alır
		} else { // id parametresi yoksa
			h.getUserByUsername(w, r) // username parametresi ile kullanıcıyı alır
		}
	case http.MethodPut: // PUT istegise kullanıcıyı günceller
		h.updateUser(w, r) // id parametresi ile kullanıcıyı günceller
	case http.MethodDelete: // DELETE istegise kullanıcıyı siler
		h.deleteUser(w, r) // id parametresi ile ku llanıcıyı siler
	default: // istek metoduna ağırlık verilmez
		w.WriteHeader(http.StatusMethodNotAllowed) // istek metoduna ağırlık verir
	}
}

// createUser kullanıcı oluşturur
func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User                       // kullanıcı nesnesi oluşturur
	err := json.NewDecoder(r.Body).Decode(&user) // json ile kullanıcıyı alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.CreateUser(&user) // kullanıcıyı veritabanına kaydeder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusCreated) // istek metoduna ağırlık verir
}

// getUserByID kullanıcıyı alır
func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request, id string) {
	userID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	user, err := h.service.GetUserByID(userID) // id parametresi ile kullanıcıyı alır
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(user) // kullanıcıyı json ile döndürür
}

// getUserByUsername kullanıcıyı alır
func (h *UserHandler) getUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username") // username parametresini alır
	if username == "" {                       // username parametresi yoksa
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	user, err := h.service.GetUserByUsername(username) // username parametresi ile kullanıcıyı alır
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // istek metoduna ağırlık verir
		return
	}
	json.NewEncoder(w).Encode(user) // kullanıcıyı json ile döndürür
}

// updateUser kullanıcıyı günceller
func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User                       // kullanıcı nesnesi oluşturur
	err := json.NewDecoder(r.Body).Decode(&user) // json ile kullanıcıyı alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.UpdateUser(&user) // kullanıcıyı veritabanına kaydeder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusOK) // istek metoduna ağırlık verir
}

// deleteUser kullanıcıyı siler
func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // id parametresini alır
	if id == "" {                 // id parametresi yoksa
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	userID, err := uuid.Parse(id) // id parametresini alır
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // istek metoduna ağırlık verir
		return
	}
	err = h.service.DeleteUser(userID) // id parametresi ile kullanıcıyı siler
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // istek metoduna ağırlık verir
		return
	}
	w.WriteHeader(http.StatusOK) // istek metoduna ağırlık verir
}
