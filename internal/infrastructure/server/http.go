package server

import (
	"github.com/EdwinBaeza05/echo_app/internal/core/services/userservice"
	"github.com/EdwinBaeza05/echo_app/internal/infrastructure/repositories/memory"
	"github.com/labstack/echo/v4"
)

// RegisterRouter Echo
func RegisterRouter(engine *echo.Echo) {

	userRepository := memory.NewUserRepository()
	userService := userservice.NewService(userRepository)
	getEndpoint := GetPersonEndpoint(userService)

	engine.GET("person/", getEndpoint)

}
