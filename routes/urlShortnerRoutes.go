package routes

import (
	"urlshortner/handlers"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	o := e.Group("/o")
	o.GET("/checkstatus", handlers.ServerRunning)

}
