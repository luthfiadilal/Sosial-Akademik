package handlers

import (
	"akademik-backend/config"
	"akademik-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBillingsByStudent mengambil tagihan UKT/Lainnya per mahasiswa
func GetBillingsByStudent(c *gin.Context) {
	studentID := c.Param("student_id")

	var billings []models.Billing
	if err := config.DB.Where("student_id = ?", studentID).Find(&billings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data tagihan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": billings})
}

// GetPaymentsByBilling mengambil riwayat pembayaran sebuah tagihan
func GetPaymentsByBilling(c *gin.Context) {
	billingID := c.Param("billing_id")

	var payments []models.Payment
	if err := config.DB.Where("billing_id = ?", billingID).Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil riwayat pembayaran"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}
