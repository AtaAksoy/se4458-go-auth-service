package service

import (
	"errors"

	"github.com/AtaAksoy/se4458-go-auth-service/internal/model"
	"github.com/AtaAksoy/se4458-go-auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Register(req model.RegisterRequest) error {
	if s.UserRepo.ExistsByEmail(req.Email) {
		return errors.New("email already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}
	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(req model.LoginRequest) error {
	user, ok := s.UserRepo.GetByEmail(req.Email)
	if !ok {
		return errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
