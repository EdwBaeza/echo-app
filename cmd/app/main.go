package main

import (
	"log"

	"github.com/EdwBaeza/echo_app/internal/infrastructure/server"
	"github.com/labstack/echo/v4"
)

func main() {
	engine := echo.New()
	server.RegisterRouter(engine)
	log.Println("Starting...")
	log.Fatal(engine.Start(":8080"))
}
