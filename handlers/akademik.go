package handlers

import (
	"akademik-backend/config"
	"akademik-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	var courses []models.Course
	config.DB.Find(&courses)

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func GetClasses(c *gin.Context) {
	var classes []models.Class
	config.DB.Preload("Course").Find(&classes)

	c.JSON(http.StatusOK, gin.H{"data": classes})
}

/////////////////////////////////////////////
// COURSE CRUD
/////////////////////////////////////////////
func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&course)
	c.JSON(http.StatusCreated, gin.H{"data": course})
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&course)
	c.JSON(http.StatusOK, gin.H{"data": course})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Course{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}

/////////////////////////////////////////////
// CLASS CRUD
/////////////////////////////////////////////
func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&class)
	c.JSON(http.StatusCreated, gin.H{"data": class})
}

func UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	if err := config.DB.First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&class)
	c.JSON(http.StatusOK, gin.H{"data": class})
}

func DeleteClass(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Class{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted"})
}
