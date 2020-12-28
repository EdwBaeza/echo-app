package server

import (
	"net/http"

	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
	"github.com/EdwinBaeza05/echo_app/internal/core/ports"
	"github.com/labstack/echo/v4"
)

//GetUserHandler by lib echo v4
func GetUserHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		user, _ := service.Get(id)
		return c.JSON(http.StatusOK, user)
	}
}

//CreateUserHandler by echo v4
func CreateUserHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		var user domain.User
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		service.Create(user)
		return c.JSON(http.StatusOK, user)
	}
}
