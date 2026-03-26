package handlers

import (
	"akademik-backend/config"
	"akademik-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKrsByStudent(c *gin.Context) {
	studentID := c.Param("student_id")

	var krs []models.Krs
	config.DB.Preload("Student").Preload("Approver").Where("student_id = ?", studentID).Find(&krs)

	c.JSON(http.StatusOK, gin.H{"data": krs})
}

func CreateKrs(c *gin.Context) {
	var input models.Krs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	input.Status = "draft"
	config.DB.Create(&input)

	c.JSON(http.StatusCreated, gin.H{"data": input})
}
