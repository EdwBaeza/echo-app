package users

import (
	"net/http"

	"github.com/EdwBaeza/echo-app/internal/core/ports"
	"github.com/EdwBaeza/echo-app/internal/infrastructure/server/utils"
	"github.com/labstack/echo/v4"
)

//ListHandler by echo v4
func ListHandler(service ports.UsersService) func(c echo.Context) error {
	return func(c echo.Context) error {
		pageSize, pageNumber := utils.GetPageParams(c)
		userPage, _ := service.All(pageSize, pageNumber)
		userPage.SetLinks(c.Request())

		return c.JSON(http.StatusOK, userPage)
	}
}
