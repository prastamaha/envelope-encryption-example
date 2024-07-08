package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prastamaha/envelope-encryption-example/internal/model"
)

type Handler struct {
	Router   *mux.Router
	UserRepo model.UserRepository
}

func NewHandler(router *mux.Router, userRepo model.UserRepository) *Handler {
	return &Handler{
		Router:   router,
		UserRepo: userRepo,
	}
}

func (h *Handler) RegisterRoutes() {
	h.Router.Handle("/users", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(h.RegisterUser))).Methods(http.MethodPost)
	h.Router.Handle("/users/{username}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(h.GetUserByUsername))).Methods(http.MethodGet)
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse request
	user := model.User{
		Username:  r.FormValue("username"),
		Name:      r.FormValue("name"),
		Gender:    r.FormValue("gender"),
		Phone:     r.FormValue("phone"),
		Address:   r.FormValue("address"),
		Consented: r.FormValue("consented") == "true",
	}

	// Register user
	username, err := h.UserRepo.RegisterUser(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	response := map[string]string{
		"username": username,
		"message":  "User successfully registered",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	// Parse request
	username := mux.Vars(r)["username"]

	// Get user
	user, err := h.UserRepo.GetUserByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	response := map[string]interface{}{
		"username":   user.Username,
		"name":       user.Name,
		"gender":     user.Gender,
		"phone":      user.Phone,
		"address":    user.Address,
		"consented":  user.Consented,
		"created_at": user.CreatedAt,
	}

	userJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
