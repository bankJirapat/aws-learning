package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitServer() *echo.Echo {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"health": "Success",
		})
	})

	// api := e.Group("/aws")

	return e
}
