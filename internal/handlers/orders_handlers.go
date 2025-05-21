package handlers

import (
	"github.com/Dastan486/LazyWMS/internal/models" // Импорт моделей
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Настройка маршрутов для продуктов
func SetupOrderRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/orders", CreateOrder(db))
	r.GET("/orders", GetOrders(db))
	r.GET("/orders/:id", GetOrdersID(db))
	r.PUT("/orders/:id", UpdateOrders(db))
	r.DELETE("/orders/:id", DeleteOrder(db))
}

// Определите функции createProduct, getProducts и т.д.
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Orders
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&product)
		c.JSON(201, product)
	}
}

// Обработчик для получения всех продуктов
func GetOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []models.Orders
		db.Find(&orders)
		c.JSON(200, orders)
	}
}

// Обработчик для получения продукта по ID
func GetOrdersID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Orders
		if err := db.First(&order, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(200, order)
	}
}

// Обработчик для обновления продукта
func UpdateOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders models.Orders
		if err := db.First(&orders, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		if err := c.ShouldBindJSON(&orders); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Save(&orders)
		c.JSON(200, orders)
	}
}

// Обработчик для удаления продукта
func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Delete(&models.Orders{}, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(204, nil) // No content response
	}
}
