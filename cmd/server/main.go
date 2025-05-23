package main

import (
	"github.com/Dastan486/LazyWMS/internal/db"
	// "github.com/Dastan486/LazyWMS/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database := db.InitDB()

	// Автоматическая миграция структур для создания таблиц в базе данных
	db.AutoMigrate(database)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	// Обслуживать статические файлы из каталога ./static
	r.Static("/static", "./static")

	// authC := &handlers.AuthController{DB: database}
	// Настройка маршрутов
	// handlers.SetupProductRoutes(r, database)
	// handlers.SetupSupplierRoutes(r, database)
	// handlers.SetupInventoryTransaction(r, database)
	// handlers.SetupUserRoutes(authC, r, database)

	// Запуск сервера на порту 8080
	log.Println("Serving on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
