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
		var employees models.Employees
		if err := c.ShouldBindJSON(&employees); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Здесь можно добавить логику проверки существования продукта и поставщика

		db.Create(&employees)
		c.JSON(201, employees)
	}
}

// Обработчик для получения всех операций с товарами
func getInventoryTransactions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var employees []models.Employees
		db.Find(&employees)
		c.JSON(200, employees)
	}
}
