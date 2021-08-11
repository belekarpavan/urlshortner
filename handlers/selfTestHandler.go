package handlers

import (
	"net/http"
	"urlshortner/models"

	"github.com/labstack/echo"
)

func ServerRunning(c echo.Context) error {

	return c.JSON(http.StatusOK, models.Response{"Server is Running on " + c.Request().Host, ""})
}
