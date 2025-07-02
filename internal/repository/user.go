package repository

import (
	"github.com/AtaAksoy/se4458-go-auth-service/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) ExistsByEmail(email string) bool {
	var count int64
	r.DB.Model(&model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *UserRepository) Create(user model.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetByEmail(email string) (model.User, bool) {
	var user model.User
	result := r.DB.Where("email = ?", email).First(&user)
	return user, result.Error == nil
}
