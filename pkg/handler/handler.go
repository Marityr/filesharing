package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", HelloMessage)

	return e
}
