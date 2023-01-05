package middleware

import (
	"errors"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/constanta"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
)

func MustAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := c.Cookie("session")
		if token == nil {
			return c.Redirect(http.StatusFound, "/login")
		}
		tokenStr := token.Value
		claims, err := utility.JWTValidation(tokenStr)
		if err != nil {
			if errors.Is(err, constanta.TokenExpired) || errors.Is(err, constanta.TokenInvalid) {
				return c.Redirect(http.StatusFound, "/login")
			}
			statusCode := http.StatusInternalServerError
			data := pongo2.Context{
				"status_code": statusCode,
				"message":     "Internal Server Error",
			}
			return c.Render(statusCode, config.FromViews("/error"), data)
		}
		c.Set("user_claims", claims)
		return next(c)
	}
}

func WithToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := c.Cookie("session")
		if token != nil {
			tokenStr := token.Value
			claims, err := utility.JWTValidation(tokenStr)
			if err == nil {
				c.Set("username", claims.Username)
				c.Set("profile_pic", claims.ProfilePic)
			}
		}
		return next(c)
	}
}
