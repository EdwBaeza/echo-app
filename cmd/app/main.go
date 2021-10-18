package main

import (
	"log"

	"github.com/EdwBaeza/echo-app/internal/infrastructure/server"
	// _ "github.com/joho/godotenv/autoload"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	engine := echo.New()
	server.RegisterRouter(engine)
	log.Println("Starting...")
	log.Fatal(engine.Start(":8080"))
}
