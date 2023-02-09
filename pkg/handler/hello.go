package handler

import (
	"log"
	"net/http"

	"github.com/Marityr/filesharing/pkg/repository"
	"github.com/labstack/echo"
)

func HelloMessage(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		return nil
	}

	data, err := repository.GetHelloWorld(id)
	if err != nil {
		log.Println(err)
	}
	c.String(http.StatusOK, data)
	return nil
}
