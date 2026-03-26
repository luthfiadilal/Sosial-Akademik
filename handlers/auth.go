package handlers

import (
	"akademik-backend/config"
	"akademik-backend/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Helper: Generate JWT Token
func generateToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecretkey"
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token berlaku 72 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func Register(c *gin.Context) {
	var input struct {
		Name           string `json:"name" binding:"required"`
		Email          string `json:"email" binding:"required,email"`
		Password       string `json:"password" binding:"required,min=6"`
		Role           string `json:"role"`
		Npm            string `json:"npm"`
		Nidn           string `json:"nidn"`
		StudyProgramID uint   `json:"study_program_id"`
		Angkatan       string `json:"angkatan"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate if email already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Default role
	role := input.Role
	if role == "" {
		role = "mahasiswa"
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     role,
		IsActive: true,
	}

	tx := config.DB.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Buat Role Profile Spesifik
	switch role {
	case "mahasiswa":
		student := models.Student{
			UserID:         user.ID,
			NPM:            input.Npm,
			StudyProgramID: input.StudyProgramID,
			Angkatan:       input.Angkatan,
			Status:         "aktif",
		}
		if student.NPM == "" {
			student.NPM = "NPM-" + strconv.FormatInt(time.Now().Unix(), 10)
		}
		if student.StudyProgramID == 0 {
			student.StudyProgramID = 1
		}
		if student.Angkatan == "" {
			student.Angkatan = strconv.Itoa(time.Now().Year())
		}
		if err := tx.Create(&student).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student profile: " + err.Error()})
			return
		}

	case "dosen":
		lecturer := models.Lecturer{
			UserID: user.ID,
			NIDN:   input.Nidn,
		}
		if lecturer.NIDN == "" {
			lecturer.NIDN = "NIDN-" + strconv.FormatInt(time.Now().Unix(), 10)
		}
		if input.StudyProgramID != 0 {
			spid := input.StudyProgramID
			lecturer.StudyProgramID = &spid
		}
		if err := tx.Create(&lecturer).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lecturer profile: " + err.Error()})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create JWT token
	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func GetProfile(c *gin.Context) {
	// Retrieve user_id from context (set by AuthMiddleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	profileData := gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"is_active":  user.IsActive,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	// Load specific role profile
	switch user.Role {
	case "mahasiswa":
		var student models.Student
		config.DB.Where("user_id = ?", user.ID).First(&student)
		profileData["student_details"] = student
	case "dosen":
		var lecturer models.Lecturer
		config.DB.Where("user_id = ?", user.ID).First(&lecturer)
		profileData["lecturer_details"] = lecturer
	}

	c.JSON(http.StatusOK, gin.H{
		"data": profileData,
	})
}

func EditProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		Npm            string `json:"npm"`
		Nidn           string `json:"nidn"`
		StudyProgramID uint   `json:"study_program_id"`
		Angkatan       string `json:"angkatan"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user base profile"})
		return
	}

	// Memperbarui profil role spesifik
	switch user.Role {
	case "mahasiswa":
		var student models.Student
		if err := config.DB.Where("user_id = ?", user.ID).First(&student).Error; err == nil {
			if input.Npm != "" {
				student.NPM = input.Npm
			}
			if input.StudyProgramID != 0 {
				student.StudyProgramID = input.StudyProgramID
			}
			if input.Angkatan != "" {
				student.Angkatan = input.Angkatan
			}
			config.DB.Save(&student)
		}
	case "dosen":
		var lecturer models.Lecturer
		if err := config.DB.Where("user_id = ?", user.ID).First(&lecturer).Error; err == nil {
			if input.Nidn != "" {
				lecturer.NIDN = input.Nidn
			}
			if input.StudyProgramID != 0 {
				spid := input.StudyProgramID
				lecturer.StudyProgramID = &spid
			}
			config.DB.Save(&lecturer)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "data": gin.H{"name": user.Name, "email": user.Email}})
}
