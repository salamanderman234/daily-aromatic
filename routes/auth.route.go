package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type authRoute struct {
	mustGuestGroup *echo.Group
	handler        domain.AuthHandler
}

func NewAuthRoute(g *echo.Group, h domain.AuthHandler) domain.Route {
	return &authRoute{
		mustGuestGroup: g,
		handler:        h,
	}
}

func (u *authRoute) Register() {
	// add path
	u.mustGuestGroup.POST("login", u.handler.LoginProcess)
	u.mustGuestGroup.POST("register", u.handler.RegisterProcess)
}
