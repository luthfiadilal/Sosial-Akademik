package main

import (
	"log"
	"net/http"

	"akademik-backend/config"

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

	// Group Route untuk modul
	api := r.Group("/api/v1")
	{
		// Akademik Routes
		akademik := api.Group("/akademik")
		{
			akademik.GET("/krs", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "KRS Data"}) })
			akademik.GET("/transkrip", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Transkrip Data"}) })
		}

		// Keuangan Routes
		keuangan := api.Group("/keuangan")
		{
			keuangan.GET("/tagihan", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Tagihan UKT"}) })
		}

		// Administratif Routes
		admin := api.Group("/administratif")
		{
			admin.GET("/pengajuan", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Daftar Pengajuan"}) })
		}

		// Sosial Routes
		sosial := api.Group("/sosial")
		{
			sosial.GET("/feed", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Sosial Feed Data"}) })
		}
	}

	log.Println("Server running on port 8080...")
	r.Run(":8080")
}
