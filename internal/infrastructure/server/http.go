package server

import (
	"github.com/EdwinBaeza05/echo_app/internal/core/services/userservice"
	"github.com/EdwinBaeza05/echo_app/internal/infrastructure/repositories/nosql"
	"github.com/labstack/echo/v4"
)

// RegisterRouter Echo
func RegisterRouter(engine *echo.Echo) {

	userRepository := nosql.NewUserRepository()
	userRepository.GetClient()
	userService := userservice.NewService(userRepository)

	getUserHandler := GetUserHandler(userService)
	createUserHandler := CreateUserHandler(userService)

	engine.GET("users/:id", getUserHandler)
	engine.POST("users/", createUserHandler)
}
