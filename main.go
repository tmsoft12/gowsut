package main

import (
	"log"
	"os"
	"tm/database"
	"tm/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	database.InitDatabase()
	// .env faýlyny ýükle
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ýalňyşlyk: .env faýlyny okap bolmady")
	}

	// IP we port üýtgeýänlerini al
	ip := os.Getenv("HOST")
	if ip == "" {
		ip = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	// CORS konfigurasiýasy
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
	}))

	// Router'lary işe girizmek
	routers.InitRouters(app)

	// Serveri işlemek
	log.Printf("Serwer %s:%s-da işläp başlady", ip, port)
	if err := app.Listen(ip + ":" + port); err != nil {
		log.Fatalf("Serweri işe girizip bolmady: %v", err)
	}
}
