package server

import (
	"log"
	"net/http"

	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
	"github.com/EdwinBaeza05/echo_app/internal/core/ports"
	"github.com/labstack/echo/v4"
)

//GetUserHandler by echo v4
func GetUserHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		log.Println("Id param: ", id)
		user, _ := service.Get(id)
		return c.JSON(http.StatusOK, user)
	}
}

//CreateUserHandler by echo v4
func CreateUserHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		var user, createdUser domain.User
		var createdError error

		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		createdUser, createdError = service.Create(user)

		if createdError != nil {
			return c.String(http.StatusInternalServerError, createdError.Error())
		}
		return c.JSON(http.StatusOK, createdUser)
	}
}
