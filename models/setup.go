package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("gocms.db"), &gorm.Config{})
	if err != nil {
		panic("Gagal konek ke database! Master apakan komputernya?!")
	}

	// Auto Migrate: Membuat tabel otomatis
	database.AutoMigrate(&User{}, &Post{})

	DB = database
}