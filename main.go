package main

import (
	"gocms/controllers"
	"gocms/middleware"
	"gocms/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// Setup CORS (PENTING BIAR ASTRO BISA AKSES)
	// Izinkan frontend Master (default Astro port 4321)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321", "https://falhafizh.vercel.app"} 
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Serve folder uploads sebagai file statis
	r.Static("/uploads", "./uploads")

	// Routes
	api := r.Group("/api")
	{
		// Public Routes
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register) // Hati-hati, amankan ini nanti!
		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)

		// Protected Routes (Harus pakai Token)
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/posts", controllers.CreatePost)
			protected.POST("/upload", controllers.UploadImage)
			// Tambahkan PUT/DELETE di sini
		}
	}

	// Jalankan di port 8080
	r.Run(":8080")
}