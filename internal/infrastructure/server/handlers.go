package server

import (
	"net/http"

	"github.com/EdwinBaeza05/echo_app/internal/core/ports"
	"github.com/labstack/echo/v4"
)

// GetPersonEndpoint by echo v4
func GetPersonEndpoint(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
