package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetJSON(c echo.Context) error {
	// Placeholder for API response
	response := map[string]string{
		"dummyString":              "simple string",
		"https://kedalab.com/test": "url",
	}

	return c.JSON(http.StatusOK, response)
}
