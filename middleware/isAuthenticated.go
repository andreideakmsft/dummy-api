package middleware

import "github.com/labstack/echo/v4"

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement auth
		return next(c)
	}
}
