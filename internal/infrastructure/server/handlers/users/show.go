package users

import (
	"net/http"

	"github.com/EdwBaeza/echo-app/internal/core/ports"
	"github.com/labstack/echo/v4"
)

//ShowHandler by echo v4
func ShowHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		user, err := service.Find(id)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, user)
	}
}
