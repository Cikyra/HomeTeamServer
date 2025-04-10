package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"HomeTeamServer/models"
	"github.com/google/uuid"
)

type UserHandler struct {
	// TODO: Maybe implement a UserRepository? Use a DI Framework??
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getUsers(w, r)
	case "POST":
		h.createUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{ID: uuid.New(), Name: "Gojo Satoru", Email: "gojo.satoru@gmail.com", PhotoUrls: []string{}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New(), Name: "Geto Suguru", Email: "geto.suguru@gmail.com", PhotoUrls: []string{}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	decodeError := json.NewDecoder(r.Body).Decode(&newUser)
	if decodeError != nil {
		http.Error(w, decodeError.Error(), http.StatusBadRequest)
		return
	}
	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	fmt.Printf("newUser: %+v\n", newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encodeError := json.NewEncoder(w).Encode(newUser)
	if encodeError != nil {
		http.Error(w, encodeError.Error(), http.StatusInternalServerError)
		return
	}
}
