package routes

import (
	"urlshortner/handlers"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	o := e.Group("/o")
	e.GET("/:alias", handlers.RedirectURLHandler)
	o.GET("/checkstatus", handlers.ServerRunning)
	o.POST("/create", handlers.CreateShortURLHandler)
	o.POST("/get", handlers.GetOriginalURLHandler)

}
