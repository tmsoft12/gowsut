package main

import (
	"tm/database"
	"tm/routers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.InitDatabase()
	app := fiber.New()

	// CORS middleware'i ekleme
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",                               // İzin verilen originler
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS", // İzin verilen HTTP yöntemleri
		AllowHeaders:     "Origin, Content-Type, Accept",    // İzin verilen başlıklar
		AllowCredentials: false,                             // Kredansiyel izinleri (örn. cookie'ler)
	}))

	routers.InitRouters(app)
	app.Listen(":3000")
}
