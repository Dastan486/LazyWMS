package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Product struct represents the products table
type Product struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Supplier struct represents the suppliers table
type Supplier struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
}

var db *gorm.DB

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Serve static files from the ./static directory
	r.Static("/static", "./static")

	// Routes for products
	r.POST("/products", createProduct)
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	// Routes for suppliers
	r.POST("/suppliers", createSupplier)
	r.GET("/suppliers", getSuppliers)

	return r
}

// Handlers for Products
func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Create(&product)
	c.JSON(201, product)
}

func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)

	c.JSON(200, products)
}

func getProduct(c *gin.Context) {
	var product Product
	if err := db.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func updateProduct(c *gin.Context) {
	var product Product
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

func deleteProduct(c *gin.Context) {
	if err := db.Delete(&Product{}, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(204, nil) // No content response
}

// Handlers for Suppliers
func createSupplier(c *gin.Context) {
	var supplier Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Create(&supplier)
	c.JSON(201, supplier)
}

func getSuppliers(c *gin.Context) {
	var suppliers []Supplier
	db.Find(&suppliers)

	c.JSON(200, suppliers)
}

func main() {
	var err error

	dsn := "host=localhost user=postgres password=Tristan dbname=lazydb port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return
	}

	// Auto migrate all structs to create tables in the database
	db.AutoMigrate(&Product{}, &Supplier{})

	r := setupRouter()

	// Start the server on port 8080
	log.Println("Serving on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
