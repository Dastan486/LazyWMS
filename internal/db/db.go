package db

import (
	"github.com/Dastan486/LazyManager/internal/models" // Импорт моделей
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func InitDB() *gorm.DB {
	var err error
	dsn := "host=localhost user=postgres password=Tristan dbname=lazydb port=5432 sslmode=disable"

	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil
	}

	return database
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Product{}, &models.Supplier{}) // Используйте модели
	db.AutoMigrate(&models.InventoryTransaction{}, &models.ProductSupplier{})
}
