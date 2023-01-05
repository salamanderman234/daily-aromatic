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
	utility "github.com/salamanderman234/daily-aromatic/utilities"
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
	isValid, credError := creds.Check()
	if !isValid {
		for _, cookie := range credError.ErrorCookies {
			c.SetCookie(cookie)
		}
		return c.Redirect(statusCode, "/login")
	}
	// login using service
	token, err := a.serv.Login(c.Request().Context(), creds.Username, creds.Password)
	// if username or password is incorrect
	if err != nil {
		if errors.Is(err, constanta.UserDataNotFoundWithCreds) {
			cookiesName := []string{"user", "pass"}
			errCookieList := utility.SameErrorCookieGen(cookiesName, "error", constanta.UserDataNotFoundWithCreds.Error())
			for _, cookie := range errCookieList {
				c.SetCookie(cookie)
			}
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
	// init
	statusCode := http.StatusFound
	data := pongo2.Context{}
	creds := entity.RegisterCred{}
	// bind
	if err := c.Bind(&creds); err != nil {
		statusCode = http.StatusBadRequest
		data["err"] = "Bad Request"
		data["status_code"] = statusCode
		return c.Render(statusCode, config.FromViews("/error.html"), data)
	}
	// perform cred check
	isValid, registerCredErr := creds.Check()
	if !isValid {
		for _, cookie := range registerCredErr.ErrorCookies {
			c.SetCookie(cookie)
		}
		return c.Redirect(statusCode, "/register")
	}
	// check if username already exists
	err := a.serv.IsUsernameAlreadyExists(c.Request().Context(), creds.Username)
	if err != nil {
		if errors.Is(err, constanta.UsernameAlreadyExists) {
			cookie := &http.Cookie{
				Name:  "user_error",
				Value: err.Error(),
				Path:  "/",
			}
			c.SetCookie(cookie)
			// set cookie
			return c.Redirect(statusCode, "/register")
		}
		statusCode = http.StatusInternalServerError
		data["err"] = "Internal Server Error"
		data["status_code"] = statusCode
		return c.Render(statusCode, config.FromViews("/error.html"), data)
	}
	// register user using service
	token, err := a.serv.Register(c.Request().Context(), creds.Username, creds.Password)
	if err != nil {
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
