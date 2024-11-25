package handlers

import (
	"net/http"

	"github.com/Dastan486/LazyWMS/internal/middleware" // Импорт моделей
	"github.com/Dastan486/LazyWMS/internal/models"     // Импорт моделей
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up the routes for user registration and login
func SetupUserRoutes(authC *AuthController, r *gin.Engine, db *gorm.DB) {
	r.GET("/login", getLogin)
	r.GET("/register", getRegister)
	r.POST("/register", authC.Register)
	r.POST("/login", authC.Register)
	r.GET("/protected-route", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a protected route!"})
	})

	r.GET("/profile", getProfile)

}
func getProfile(c *gin.Context) {

	c.HTML(http.StatusOK, "profile.html", nil)
}
func getLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}

func getRegister(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", nil)
}

type AuthController struct {
	DB *gorm.DB
}

// Register обрабатывает регистрацию пользователя
func (ctrl *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Check if the email already exists
	var existingUser models.User
	if err := ctrl.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Create the user if the email does not exist
	if err := ctrl.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User registration failed"})
		return
	}

	// Generate token after successful registration
	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "token": token})
}
func (ctrl *AuthController) Login(c *gin.Context) {
	var user models.User
	var input models.User // Structure for login (email and password)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email
	if err := ctrl.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token after successful login
	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
