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

func GetOriginalURLHandler(c echo.Context) error {
	request := models.RequestData{}
	c.Bind(&request)
	response := models.Response{}
	url, err := services.GetOriginalURLService(request)
	if err != nil {
		response.Error = "ERROR_FAILED_TO_GET_ORIGINAL_URL"
		return c.JSON(http.StatusInternalServerError, response)
	}
	response.Result = url
	return c.JSON(http.StatusInternalServerError, response)
}

func RedirectURLHandler(c echo.Context) error {
	response := models.Response{}
	alias := c.Param("alias")
	if alias == "" {
		response.Error = "Enter valid url"
		c.JSON(http.StatusInternalServerError, response)
	}

	url, err := services.GetRedirectURLService(alias)
	if err != nil {
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError, response)
	}
	return c.Redirect(301, url)
}
