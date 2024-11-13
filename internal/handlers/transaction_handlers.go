package handlers

import (
	"github.com/Dastan486/LazyWMS/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupInventoryTransaction(r *gin.Engine, db *gorm.DB) {
	// Маршруты для учета операций с товарами
	r.POST("/inventory/transactions", createInventoryTransaction(db))
	r.GET("/inventory/transactions", getInventoryTransactions(db))

}

// Обработчик для создания операции с товаром
func createInventoryTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var transaction models.InventoryTransaction
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Здесь можно добавить логику проверки существования продукта и поставщика

		db.Create(&transaction)
		c.JSON(201, transaction)
	}
}

// Обработчик для получения всех операций с товарами
func getInventoryTransactions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var transactions []models.InventoryTransaction
		db.Find(&transactions)
		c.JSON(200, transactions)
	}
}
