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

type userHandler struct {
	serv domain.UserService
}

func NewUserHandler(s domain.UserService) domain.UserHandler {
	return &userHandler{
		serv: s,
	}
}

func (u *userHandler) UpdateProfile(c echo.Context) error {
	// get id
	id := c.Get("id")
	username := c.Get("username")
	if id == nil || username == nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusUnauthorized)
		return c.Render(status, config.FromViews(fileName), data)
	}
	idUint, ok := id.(uint)
	if !ok {
		status, data, fileName := utility.ErrorPageFactory(http.StatusUnauthorized)
		return c.Render(status, config.FromViews(fileName), data)
	}
	updateField := entity.User{}
	if err := c.Bind(&updateField); err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
	}
	isValid, updateErr := updateField.Validate()
	if !isValid {
		for _, cookie := range updateErr.ErrorCookies {
			c.SetCookie(cookie)
		}
		return c.Redirect(http.StatusFound, "/profile")
	}
	if username.(string) != updateField.Username {
		err := u.serv.IsUsernameAlreadyExists(c.Request().Context(), updateField.Username)
		if err != nil {
			if errors.Is(err, variable.ErrMustUniqueUsername) {
				cookie := &http.Cookie{
					Name:  variable.UsernameErrCookie,
					Value: err.Error(),
					Path:  "/",
				}
				c.SetCookie(cookie)
				// set cookie
				return c.Redirect(http.StatusFound, "/profile")
			}
			status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
			return c.Render(status, config.FromViews(fileName), data)
		}
	}
	jwt, err := u.serv.UpdateUser(c.Request().Context(), idUint, updateField)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// set new cookie
	session_cookie := &http.Cookie{
		Name:    variable.SessionCookie,
		Value:   jwt,
		Expires: time.Now().Add(variable.TokenAndSessionValidityRange * time.Hour),
		Path:    "/",
	}
	c.SetCookie(session_cookie)
	return c.Redirect(http.StatusFound, "/profile")
}

func (u *userHandler) UpdateProfilePic(c echo.Context) error {
	avatar := entity.UserAvatar{}
	if err := c.Bind(&avatar); err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
	}
	id := c.Get("id")
	idUint, ok := id.(uint)
	if !ok {
		status, data, fileName := utility.ErrorPageFactory(http.StatusUnauthorized)
		return c.Render(status, config.FromViews(fileName), data)
	}
	user := entity.User{Avatar: avatar.Avatar}
	jwt, err := u.serv.UpdateUser(c.Request().Context(), idUint, user)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// set new cookie
	session_cookie := &http.Cookie{
		Name:    variable.SessionCookie,
		Value:   jwt,
		Expires: time.Now().Add(variable.TokenAndSessionValidityRange * time.Hour),
		Path:    "/",
	}
	c.SetCookie(session_cookie)
	return c.Redirect(http.StatusFound, "/profile")
}
