package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type userViewRoute struct {
	router  *echo.Echo
	handler domain.UserViewHandler
}

func NewUserViewRoute(r *echo.Echo, h domain.UserViewHandler) domain.Route {
	return &userViewRoute{
		router:  r,
		handler: h,
	}
}

func (u *userViewRoute) Register() {
	group := u.router.Group("/")

	// add path
	group.GET("", u.handler.PageLanding)
	group.GET("login", u.handler.PageLogin)
	group.GET("register", u.handler.PageRegister)
}
