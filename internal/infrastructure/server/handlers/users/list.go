package users

import (
	"net/http"

	"github.com/EdwBaeza/echo-app/internal/core/ports"
	"github.com/labstack/echo/v4"
)

//ListHandler by echo v4
func ListHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		users, _ := service.All()
		return c.JSON(http.StatusOK, users)
	}
}
