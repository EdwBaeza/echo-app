package users

import (
	"log"
	"net/http"

	"github.com/EdwBaeza/echo-app/internal/core/ports"
	"github.com/labstack/echo/v4"
)

//ShowHandler by echo v4
func ShowHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		log.Println("Id param: ", id)
		user, _ := service.Find(id)
		return c.JSON(http.StatusOK, user)
	}
}
