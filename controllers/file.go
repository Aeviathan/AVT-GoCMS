package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Mana gambarnya?!"})
		return
	}

	// Ganti nama file biar unik
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	path := "uploads/" + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(500, gin.H{"error": "Gagal simpan gambar. Server lelah."})
		return
	}

	// Balikin URL lengkap (sesuaikan domain nanti saat deploy)
	fullURL := "http://localhost:8080/uploads/" + filename
	c.JSON(200, gin.H{"url": fullURL})
}