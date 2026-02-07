package controllers

import (
	"gocms/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Public: Ambil semua post
func GetPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)
	c.JSON(200, gin.H{"data": posts})
}

// Public: Ambil detail post by ID
func GetPost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": "Postingan hilang ditelan bumi!"})
		return
	}
	c.JSON(200, gin.H{"data": post})
}

// Protected: Buat post baru
func CreatePost(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Buat slug sederhana
	input.Slug = strings.ReplaceAll(strings.ToLower(input.Title), " ", "-")
	
	models.DB.Create(&input)
	c.JSON(200, gin.H{"data": input})
}