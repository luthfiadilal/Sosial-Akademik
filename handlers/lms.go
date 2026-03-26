package handlers

import (
	"akademik-backend/config"
	"akademik-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List Materials by Class
func GetCourseMaterials(c *gin.Context) {
	classID := c.Param("class_id")

	var materials []models.CourseMaterial
	if err := config.DB.Where("class_id = ?", classID).Find(&materials).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch materials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": materials})
}

// List Assignments by Class
func GetAssignments(c *gin.Context) {
	classID := c.Param("class_id")

	var assignments []models.Assignment
	if err := config.DB.Where("class_id = ?", classID).Find(&assignments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch assignments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assignments})
}

// Get Quizzes by Class
func GetQuizzes(c *gin.Context) {
	classID := c.Param("class_id")

	var quizzes []models.Quiz
	if err := config.DB.Where("class_id = ?", classID).Find(&quizzes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quizzes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": quizzes})
}

// Get Forums by Class
func GetForums(c *gin.Context) {
	classID := c.Param("class_id")

	var forums []models.Forum
	if err := config.DB.Where("class_id = ?", classID).Find(&forums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch forums"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": forums})
}
