package controllers

import (
	"gocms/middleware"
	"gocms/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register admin baru (Hapus endpoint ini kalau sudah punya admin!)
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Username: input.Username, Password: string(hashedPassword)}
	models.DB.Create(&user)

	c.JSON(200, gin.H{"message": "Admin berhasil dibuat!"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Input salah, baka!"})
		return
	}

	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Username tidak ditemukan!"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Password salah!"})
		return
	}

	token, _ := middleware.GenerateToken(user.ID)
	c.JSON(200, gin.H{"token": token})
}