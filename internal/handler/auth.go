package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/AtaAksoy/se4458-go-auth-service/internal/model"
	"github.com/AtaAksoy/se4458-go-auth-service/internal/service"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecret"
	}
	return []byte(secret)
}

func GenerateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Register godoc
// @Summary Register new user
// @Description Registers a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   registerRequest body model.RegisterRequest true "Register Request"
// @Success 201 {object} model.AuthResponse
// @Failure 400 {object} model.AuthResponse
// @Router /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, model.AuthResponse{Status: "error", Error: "Invalid request"})
		return
	}
	err := h.AuthService.Register(req)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, model.AuthResponse{Status: "error", Error: err.Error()})
		return
	}
	user, _ := h.AuthService.UserRepo.GetByEmail(req.Email)
	token, _ := GenerateJWT(user.ID, user.Email)
	respondJSON(w, http.StatusCreated, model.AuthResponse{
		Status: "success",
		Token:  token,
		User:   &model.UserPublic{ID: user.ID, Email: user.Email},
	})
}

// Login godoc
// @Summary Login user
// @Description Logs in a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   loginRequest body model.LoginRequest true "Login Request"
// @Success 200 {object} model.AuthResponse
// @Failure 401 {object} model.AuthResponse
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, model.AuthResponse{Status: "error", Error: "Invalid request"})
		return
	}
	err := h.AuthService.Login(req)
	if err != nil {
		respondJSON(w, http.StatusUnauthorized, model.AuthResponse{Status: "error", Error: err.Error()})
		return
	}
	user, _ := h.AuthService.UserRepo.GetByEmail(req.Email)
	token, _ := GenerateJWT(user.ID, user.Email)
	respondJSON(w, http.StatusOK, model.AuthResponse{
		Status: "success",
		Token:  token,
		User:   &model.UserPublic{ID: user.ID, Email: user.Email},
	})
}
