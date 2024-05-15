package main

import (
	"log"
	"os"

	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/database"
	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	env := godotenv.Load()
	if env != nil {
		panic("Failed to load env file")
	} else {
		log.Println("Connect seccessfully")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.SetUp(app)
	app.Listen(":" + port)
}
