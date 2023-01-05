package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
	middleware "github.com/salamanderman234/daily-aromatic/middlewares"
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
	// group := u.router.Group("/")
	groupWithToken := u.router.Group("/", middleware.WithToken)
	groupGuest := u.router.Group("/", middleware.MustGuest)
	// add path
	groupWithToken.GET("", u.handler.PageLanding)
	groupGuest.GET("login", u.handler.PageLogin)
	groupGuest.GET("register", u.handler.PageRegister)
}
