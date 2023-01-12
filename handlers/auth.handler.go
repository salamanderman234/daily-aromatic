package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

type authHandler struct {
	serv domain.AuthService
}

func NewAuthHandler(s domain.AuthService) domain.AuthHandler {
	return &authHandler{
		serv: s,
	}
}

func (a *authHandler) LoginProcess(c echo.Context) error {
	// init
	statusCode := http.StatusFound
	creds := entity.Credentials{}
	// binding form
	if err := c.Bind(&creds); err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// cred validation
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
		if errors.Is(err, variable.ErrUserCredsNotFound) {
			cookiesName := []string{variable.UsernameErrCookie, variable.PasswordErrCookie}
			errCookieList := utility.SameCookieGen(cookiesName, variable.ErrUserCredsNotFound.Error())
			for _, cookie := range errCookieList {
				c.SetCookie(cookie)
			}
			return c.Redirect(statusCode, "/login")
		}
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// if success
	session_cookie := &http.Cookie{
		Name:    variable.SessionCookie,
		Value:   token,
		Expires: time.Now().Add(variable.TokenAndSessionValidityRange * time.Hour),
		Path:    "/",
	}
	c.SetCookie(session_cookie)
	return c.Redirect(statusCode, "/")
}
func (a *authHandler) RegisterProcess(c echo.Context) error {
	// init
	statusCode := http.StatusFound
	creds := entity.RegisterCred{}
	// bind
	if err := c.Bind(&creds); err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
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
		if errors.Is(err, variable.ErrMustUniqueUsername) {
			cookie := &http.Cookie{
				Name:  variable.UsernameErrCookie,
				Value: err.Error(),
				Path:  "/",
			}
			c.SetCookie(cookie)
			// set cookie
			return c.Redirect(statusCode, "/register")
		}
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// register user using service
	token, err := a.serv.Register(c.Request().Context(), creds.Username, creds.Password)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// if success
	session_cookie := &http.Cookie{
		Name:    variable.SessionCookie,
		Value:   token,
		Expires: time.Now().Add(variable.TokenAndSessionValidityRange * time.Hour),
		Path:    "/",
	}
	c.SetCookie(session_cookie)
	return c.Redirect(statusCode, "/")
}

func (a *authHandler) LogOutProcess(c echo.Context) error {
	session, err := c.Cookie(variable.SessionCookie)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
	}
	session.Expires = time.Unix(0, 0)
	session.MaxAge = -1
	c.SetCookie(session)
	return c.Redirect(http.StatusFound, "/login")
}
