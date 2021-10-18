package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	DEFAULT_PAGE_SIZE   = 10
	DEFAULT_PAGE_NUMBER = 1
)

func GetPageParams(c echo.Context) (pageSize int, pageNumber int) {
	pageSizeParam := c.QueryParam("pageSize")
	pageNumberParam := c.QueryParam("pageNumber")

	getParam := func(param string, defaultValue int) int {
		if param == "" {
			return defaultValue
		}
		number, _ := strconv.Atoi(param)

		return number
	}

	pageSize = getParam(pageSizeParam, DEFAULT_PAGE_SIZE)
	pageNumber = getParam(pageNumberParam, DEFAULT_PAGE_NUMBER)

	return
}
