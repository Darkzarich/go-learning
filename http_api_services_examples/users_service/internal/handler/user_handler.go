package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"users-service/internal/service"
	"users-service/pkg/app_error"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// RegisterRoutes sets up the HTTP routes on a ServeMux.
func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", h.handleUsers)     // POST and GET (GET not shown)
	mux.HandleFunc("/users/", h.handleUserByID) // GET, PUT, DELETE by ID
}

func (h *UserHandler) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.svc.GetAll()
		if err != nil {
			h.writeError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var input struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}

		user, err := h.svc.Create(input.Name, input.Email)
		if err != nil {
			h.writeError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) handleUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL like /users/123
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, `{"error":"invalid user id"}`, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := h.svc.GetByID(id)
		if err != nil {
			h.writeError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	case http.MethodDelete:
		err := h.svc.DeleteByID(id)
		if err != nil {
			h.writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("OK"))
	case http.MethodPut:
		var input struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}

		user, err := h.svc.Update(id, input.Name, input.Email)
		if err != nil {
			h.writeError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// writeError translates a service error into an HTTP status and message.
func (h *UserHandler) writeError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	msg := `{"error":"internal server error"}`

	if ae, ok := err.(*app_error.AppError); ok {
		switch ae.Kind {
		case app_error.KindNotFound:
			status = http.StatusNotFound
		case app_error.KindInvalidInput:
			status = http.StatusBadRequest
		case app_error.KindAlreadyExists:
			status = http.StatusConflict
		}
		msg = `{"error":"` + ae.Message + `"}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(msg))
	log.Printf("HTTP error: %v\n", err)
}
