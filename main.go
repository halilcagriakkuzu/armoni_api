package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/akor", tirrek)

	// Start server
	e.Logger.Fatal(e.Start(":1337"))
}

func tirrek(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Akor ver!</h1>")
}
