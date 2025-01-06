package repository

import (
	"phishing-platform-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// Crear un nuevo usuario
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

// Buscar un usuario por su email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Listar todos los usuarios
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}
