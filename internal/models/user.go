package models

import "gorm.io/gorm"

// User representa un usuario en el sistema
type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`      // Nombre del usuario
	Email    string `gorm:"uniqueIndex;not null"`   // Email único
	Password string `gorm:"size:255"`               // Contraseña cifrada
	Role     string `gorm:"size:20;default:'user'"` // Rol del usuario ("admin" o "user")
	APIKey   string `gorm:"size:255"`               // API Key cifrada para GoPhish
}
