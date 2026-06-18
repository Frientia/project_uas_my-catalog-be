package main

import (
	"log"
	"os"

	"github.com/Frientia/my-firebase-backend/config"
	"github.com/Frientia/my-firebase-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables dari .env file
	if err := godotenv.Load("C:/laragon1/www/my-firebase-backend/.env"); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
	}
	// 2. Inisialisasi Firebase Admin SDK
	config.InitFirebase()
	// 3. Inisialisasi database + AutoMigrate
	config.InitDatabase()
	// 4. Setup Gin router dengan semua routes
	router := routes.SetupRouter()
	// 5. Jalankan server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server berjalan di http://localhost:8080:%s", port)
	log.Printf("Health check: http://localhost:8080:%s/v1/health", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
