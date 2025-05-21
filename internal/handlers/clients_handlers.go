package handlers

import (
	"github.com/Dastan486/LazyWMS/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupClients(r *gin.Engine, db *gorm.DB) {
	r.POST("/", CreateClient(db))
	r.GET("/suppliers", GetClients(db))
}

func CreateClient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var clinent models.Clients
		if err := c.ShouldBindJSON(&clinent); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&clinent)
		c.JSON(201, clinent)
	}
}
func GetClients(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var clinents []models.Clients
		db.Find(&clinents)
		c.JSON(200, clinents)
	}
}
