package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AtaAksoy/se4458-go-auth-service/internal/model"
	"github.com/AtaAksoy/se4458-go-auth-service/internal/service"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

// Register godoc
// @Summary Register new user
// @Description Registers a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   registerRequest body model.RegisterRequest true "Register Request"
// @Success 201 {string} string "created"
// @Failure 400 {string} string "error"
// @Router /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	err := h.AuthService.Register(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Login godoc
// @Summary Login user
// @Description Logs in a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   loginRequest body model.LoginRequest true "Login Request"
// @Success 200 {string} string "ok"
// @Failure 401 {string} string "unauthorized"
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	err := h.AuthService.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
