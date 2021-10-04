package server

import (
	"github.com/EdwBaeza/echo-app/internal/core/services/user_service"
	"github.com/EdwBaeza/echo-app/internal/infrastructure/repositories/nosql"
	"github.com/EdwBaeza/echo-app/internal/infrastructure/server/handlers/users"
	"github.com/labstack/echo/v4"
)

// RegisterRouter Echo
func RegisterRouter(engine *echo.Echo) {

	userRepository := nosql.NewUserRepository()
	userRepository.GetClient()
	userService := user_service.NewService(userRepository)

	showUserHandler := users.ShowHandler(userService)
	listUsersHandler := users.ListHandler(userService)
	createUserHandler := users.CreateHandler(userService)

	engine.GET("users/:id", showUserHandler)
	engine.GET("users/", listUsersHandler)
	engine.POST("users/", createUserHandler)
}
