package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
	variable "github.com/salamanderman234/daily-aromatic/vars"
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
			if errors.Is(err, variable.ErrTokenExpired) || errors.Is(err, variable.ErrTokenNotValid) {
				return c.Redirect(http.StatusFound, "/login")
			}
			statusCode, data, filename := utility.ErrorPageFactory(http.StatusInternalServerError)
			return c.Render(statusCode, config.FromViews(filename), data)
		}
		c.Set("id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("profile_pic", claims.ProfilePic)
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
