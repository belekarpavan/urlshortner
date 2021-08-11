package main

import (
	"urlshortner/routes"
	"urlshortner/utils"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes.Init(e)
	e.Start(utils.Port)
}
