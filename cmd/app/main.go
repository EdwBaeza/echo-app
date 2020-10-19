package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/EdwinBaeza05/echo_app/internal/infrastructure/server"
)

func main() {
	engine := echo.New()
	server.RegisterRouter(engine)
	log.Fatal(engine.Start(":8080"))
}
