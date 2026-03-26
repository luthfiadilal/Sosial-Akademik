package main

import (
	"log"
	"net/http"

	"akademik-backend/config"
	"akademik-backend/handlers"
	"akademik-backend/routes"
	"akademik-backend/ws"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan environment OS bawaan")
	}

	// Initialize Database Connection
	config.ConnectDatabase()

	// Initialize WebSocket Hub
	handlers.AppHub = ws.NewHub()
	go handlers.AppHub.Run()

	r := gin.Default()

	// CORS Setup (Sederhana untuk development)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Rute dasar (Health Check)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Welcome to Akademik Backend API",
			"version": "1.0.0",
		})
	})

	// Panggil Modular Routes
	routes.SetupRoutes(r)

	log.Println("Server running on port 8080...")
	r.Run(":8080")
}
