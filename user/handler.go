package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	name := strings.TrimSpace(input.Name)
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	_, err = h.service.CreateUser(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	if strings.TrimSpace(userID) == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}
	intUserID, err := strconv.Atoi(userID)
	if err != nil || intUserID <= 0 {
		http.Error(w, "user id is invalid", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(intUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {

	users := h.service.GetAllUsers()

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	if strings.TrimSpace(userID) == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}
