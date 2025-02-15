package main

import (
	"log"

	"github.com/PatriciaDelCastillo/Rest-Api-Goland/database"
	"github.com/PatriciaDelCastillo/Rest-Api-Goland/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()
	app:= fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",	
		AllowCredentials: true,	
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}