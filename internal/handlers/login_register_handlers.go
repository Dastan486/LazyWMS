package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// SetupUserRoutes sets up the routes for user registration and login
func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/login", getLogin)
	r.GET("/register", getRegister)

}

func getLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}

func getRegister(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", nil)
}

