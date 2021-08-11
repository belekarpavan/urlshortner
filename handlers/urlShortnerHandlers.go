package handlers

import (
	"net/http"
	"urlshortner/models"

	"urlshortner/services"

	"github.com/labstack/echo"
)

func CreateShortURLHandler(c echo.Context) error {
	request := models.RequestData{}
	c.Bind(&request)
	response := models.Response{}
	url, err := services.CreateShortURLService(request)
	if err != nil {
		response.Error = "ERROR_FAILED_TO_CREATE_SHORT_URL"
		return c.JSON(http.StatusInternalServerError, response)
	}
	response.Result = url
	return c.JSON(http.StatusInternalServerError, response)
}
