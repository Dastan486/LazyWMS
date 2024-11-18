package db

import (
	"github.com/Dastan486/LazyWMS/internal/models" // Импорт моделей
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func InitDB() *gorm.DB {
	var err error
	dsn := "root:Tristan@tcp(localhost:3306)/lazydb?charset=utf8&parseTime=True&loc=Local"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
