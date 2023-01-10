package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
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
	if id == nil {
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
	err := u.serv.UpdateUser(c.Request().Context(), idUint, updateField)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	return c.Redirect(http.StatusFound, "/profile")
}
