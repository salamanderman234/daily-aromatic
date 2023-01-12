package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type userViewRoute struct {
	withTokenGroup *echo.Group
	mustAuthGroup  *echo.Group
	mustGuestGroup *echo.Group
	handler        domain.UserViewHandler
}

func NewUserViewRoute(withTokenGroup *echo.Group,
	mustAuthGroup *echo.Group,
	mustGuestGroup *echo.Group,
	h domain.UserViewHandler,
) domain.Route {
	return &userViewRoute{
		withTokenGroup: withTokenGroup,
		mustAuthGroup:  mustAuthGroup,
		mustGuestGroup: mustGuestGroup,
		handler:        h,
	}
}

func (u *userViewRoute) Register() {
	// add path
	u.withTokenGroup.GET("", u.handler.PageLanding)
	u.mustAuthGroup.GET("profile", u.handler.PageUserProfile)
	u.withTokenGroup.GET("profile/:username", u.handler.PageDiffUserProfile)
	u.withTokenGroup.GET("products", u.handler.PageProductSearch)
	u.withTokenGroup.GET("products/:id", u.handler.ProductDetailPage)
	u.mustGuestGroup.GET("login", u.handler.PageLogin)
	u.mustGuestGroup.GET("register", u.handler.PageRegister)
	u.mustAuthGroup.GET("reviews/add", u.handler.NewReviewPage)
}
