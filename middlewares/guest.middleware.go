package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
)

func MustGuest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("session")
		if cookie == nil {
			return next(c)
		} else {
			tokenStr := cookie.Value
			_, err := utility.JWTValidation(tokenStr)
			if err != nil {
				return next(c)
			}
		}
		return c.Redirect(http.StatusFound, "/")
	}
}
