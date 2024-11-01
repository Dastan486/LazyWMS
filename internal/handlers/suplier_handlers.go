package handlers

import (
	"github.com/Dastan486/LazyWMS/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupSupplierRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/suppliers", createSupplier(db))
	r.GET("/suppliers", getSuppliers(db))
}

func createSupplier(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplier models.Supplier
		if err := c.ShouldBindJSON(&supplier); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&supplier)
		c.JSON(201, supplier)
	}
}
func getSuppliers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var suppliers []models.Supplier
		db.Find(&suppliers)
		c.JSON(200, suppliers)
	}
}

// Остальные обработчики (getSuppliers и т.д.) аналогично...
