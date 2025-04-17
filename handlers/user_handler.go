package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"HomeTeamServer/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserHandler struct {
	// TODO: Maybe implement a UserRepository? Use a DI Framework??
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewUserHandler(db *gorm.DB, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{db: db, logger: logger}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("GET %s", r.URL.Path)
	userId := r.PathValue("id")
	if userId == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var user models.User
	result := h.db.First(&user, "id = ?", userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, fmt.Sprintf("User with id(%s) not found", userId), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("GET %s", r.URL.Path)
	var users []models.User
	result := h.db.Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("POST %s", r.URL.Path)
	var newUser models.User

	decodeError := json.NewDecoder(r.Body).Decode(&newUser)
	if decodeError != nil {
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
		return
	}
	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	result := h.db.Create(&newUser)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encodeError := json.NewEncoder(w).Encode(newUser)
	if encodeError != nil {
		http.Error(w, encodeError.Error(), http.StatusInternalServerError)
		return
	}
}
