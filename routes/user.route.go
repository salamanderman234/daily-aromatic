package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type userRoute struct {
	mustAuthGroup *echo.Group
	handler       domain.UserHandler
}

func NewUserRoute(g *echo.Group, h domain.UserHandler) domain.Route {
	return &userRoute{
		mustAuthGroup: g,
		handler:       h,
	}
}

func (u *userRoute) Register() {
	// add path
	u.mustAuthGroup.POST("profile/update", u.handler.UpdateProfile)
	u.mustAuthGroup.POST("profile/avatar/update", u.handler.UpdateProfilePic)
}
