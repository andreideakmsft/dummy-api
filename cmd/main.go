package main

import (
	handlers "github.com/andreideakmsft/repro-api/handlers"
	middleware "github.com/andreideakmsft/repro-api/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Route to get the JSON data
	e.GET("/api/json", handlers.GetJSON, middleware.IsAuthenticated)

	// TODO: Route to change the JSON sent by the API

	e.Logger.Fatal(e.Start(":8080"))
}
