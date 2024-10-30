package handlers

import (
	"github.com/Dastan486/LazyManager/internal/models" // Импорт моделей
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Настройка маршрутов для продуктов
func SetupProductRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/products", createProduct(db))
	r.GET("/products", getProducts(db))
	r.GET("/products/:id", getProduct(db))
	r.PUT("/products/:id", updateProduct(db))
	r.DELETE("/products/:id", deleteProduct(db))
}

// Определите функции createProduct, getProducts и т.д.
func createProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&product)
		c.JSON(201, product)
	}
}

// Обработчик для получения всех продуктов
func getProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		db.Find(&products)
		c.JSON(200, products)
	}
}

// Обработчик для получения продукта по ID
func getProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := db.First(&product, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(200, product)
	}
}

// Обработчик для обновления продукта
func updateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := db.First(&product, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Save(&product)
		c.JSON(200, product)
	}
}

// Обработчик для удаления продукта
func deleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Delete(&models.Product{}, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(204, nil) // No content response
	}
}

