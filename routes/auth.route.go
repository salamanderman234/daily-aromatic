package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type authRoute struct {
	router  *echo.Echo
	handler domain.AuthHandler
}

func NewAuthRoute(r *echo.Echo, h domain.AuthHandler) domain.Route {
	return &authRoute{
		router:  r,
		handler: h,
	}
}

func (u *authRoute) Register() {
	group := u.router.Group("/")

	// add path
	group.POST("login", u.handler.LoginProcess)
	group.POST("register", u.handler.RegisterProcess)
}
