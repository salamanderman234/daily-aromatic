package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/constanta"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
)

type authHandler struct {
	serv domain.AuthService
}

const (
	validityRange = 24
)

func NewAuthHandler(s domain.AuthService) domain.AuthHandler {
	return &authHandler{
		serv: s,
	}
}

func (a *authHandler) LoginProcess(c echo.Context) error {
	// init
	statusCode := http.StatusFound
	data := pongo2.Context{}
	creds := entity.Credentials{}
	// binding form
	if err := c.Bind(&creds); err != nil {
		statusCode = http.StatusBadRequest
		data["err"] = "Bad Request"
		data["status_code"] = statusCode
		return c.Render(statusCode, config.FromViews("/error.html"), data)
	}
	// if username or password is empty
	if creds.Password == "" || creds.Username == "" {
		statusCode = http.StatusBadRequest
		if creds.Username == "" {
			// set err cookie if error
			cookie := &http.Cookie{
				Name:  "user_error",
				Value: constanta.EmptyCreds.Error(),
				Path:  "/",
			}
			c.SetCookie(cookie)
		}
		if creds.Password == "" {
			cookie := &http.Cookie{
				Name:  "pass_error",
				Value: constanta.EmptyCreds.Error(),
				Path:  "/",
			}
			c.SetCookie(cookie)
		}
		// set cookie
		return c.Redirect(statusCode, "/login")
	}
	token, err := a.serv.Login(c.Request().Context(), creds.Username, creds.Password)
	// if username or password is incorrect
	if err != nil {
		if errors.Is(err, constanta.UserDataNotFoundWithCreds) {
			user_error_cookie := &http.Cookie{
				Name:  "user_error",
				Value: constanta.UserDataNotFoundWithCreds.Error(),
				Path:  "/",
			}
			pass_error_cookie := &http.Cookie{
				Name:  "pass_error",
				Value: constanta.UserDataNotFoundWithCreds.Error(),
				Path:  "/",
			}
			c.SetCookie(user_error_cookie)
			c.SetCookie(pass_error_cookie)
			// set cookie
			return c.Redirect(statusCode, "/login")
		}
		statusCode = http.StatusInternalServerError
		data["err"] = "Internal Server Error"
		data["status_code"] = statusCode
		return c.Render(statusCode, config.FromViews("/error.html"), data)
	}
	// if success
	session_cookie := &http.Cookie{
		Name:    "session",
		Value:   token,
		Expires: time.Now().Add(validityRange * time.Hour),
		Path:    "/",
	}
	c.SetCookie(session_cookie)
	return c.Redirect(statusCode, "/")
}
func (a *authHandler) RegisterProcess(c echo.Context) error {
	return c.Redirect(404, "/")
}
