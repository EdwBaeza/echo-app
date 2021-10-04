package users

import (
	"net/http"

	"github.com/EdwBaeza/echo-app/internal/core/domain"
	"github.com/EdwBaeza/echo-app/internal/core/ports"
	"github.com/labstack/echo/v4"
)

//CreateHandler by echo v4
func CreateHandler(service ports.UsersService) func(c echo.Context) error {
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
