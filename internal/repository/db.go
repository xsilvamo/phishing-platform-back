package repository

import (
	"log"
	"os"
	"phishing-platform-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL no configurado en .env")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	log.Print("Conexión a la base de datos exitosa")

	// Migrar modelos
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrando modelos: %v", err)
	}

	log.Print("Migraciones completadas")

}
