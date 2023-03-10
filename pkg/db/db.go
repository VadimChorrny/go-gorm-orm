package db

import (
	"go-gorm-orm/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Ініціалізація бази даних
func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
