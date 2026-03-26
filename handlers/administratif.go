package handlers

import (
	"akademik-backend/config"
	"akademik-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotifications mengambil list notifikasi user yang login
func GetNotifications(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var notifications []models.Notification
	if err := config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil notifikasi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notifications})
}

// MarkNotificationRead menandai notifikasi telah dibaca
func MarkNotificationRead(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id") // Ambil user ID dari JWT

	var notification models.Notification
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&notification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notifikasi tidak ditemukan"})
		return
	}

	notification.IsRead = true
	config.DB.Save(&notification)

	c.JSON(http.StatusOK, gin.H{"message": "Notifikasi ditandai sudah dibaca"})
}
