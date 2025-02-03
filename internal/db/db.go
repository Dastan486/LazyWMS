package db

import (
	"github.com/Dastan486/LazyWMS/internal/models" // Импорт моделей
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"
)

var database *gorm.DB

func InitDB() *gorm.DB {
	var err error

	database, err = gorm.Open(sqlite.Open("lazy.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil
	}

	return database
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Product{}, &models.Supplier{})
	db.AutoMigrate(&models.InventoryTransaction{}, &models.ProductSupplierID{})
	db.AutoMigrate(&models.User{})
}
