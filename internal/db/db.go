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

	database, err = gorm.Open(sqlite.Open("tms.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil
	}

	return database
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Employees{},
		&models.Clients{},
		&models.Order_employees{},
		&models.Orders{},
		&models.Routes{},
		&models.Vehicles{},
	)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
}
