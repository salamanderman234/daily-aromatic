package handler

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type userViewHandler struct {
	userService         domain.UserService
	productService      domain.ProductService
	notificationService domain.NotificationService
	reviewService       domain.ReviewService
}

func NewUserViewHandler(u domain.UserService,
	p domain.ProductService,
	n domain.NotificationService,
	r domain.ReviewService,
) domain.UserViewHandler {
	return &userViewHandler{
		userService:         u,
		productService:      p,
		notificationService: n,
		reviewService:       r,
	}
}

func (u *userViewHandler) PageLanding(c echo.Context) error {
	var data pongo2.Context
	statusCode := http.StatusOK

	// reviews, err := u.reviewService.GetAllReviews()
	// if err != nil {
	// 	statusCode = http.StatusInternalServerError
	// 	data["status_code"] = statusCode
	// 	data["message"] = "Internal Server Error"

	// 	return c.Render(http.StatusInternalServerError, config.FromViews("/error.html"), data)
	// }
	statusCode = http.StatusOK
	// data["reviews"] = reviews
	data["user"] = c.Get("user")
	return c.Render(statusCode, config.FromViews("/landing.html"), data)
}
func (u *userViewHandler) PageLogin(c echo.Context) error {

}
func (u *userViewHandler) PageRegister(c echo.Context) error {

}
func (u *userViewHandler) PageProductSearch(c echo.Context) error {

}
func (u *userViewHandler) PageUserProfile(c echo.Context) error {

}
