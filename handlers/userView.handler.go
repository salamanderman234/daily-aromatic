package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
)

type userViewHandler struct {
	// userService         domain.UserService
	productService domain.ProductService
	// notificationService domain.NotificationService
	reviewService domain.ReviewService
}

func NewUserViewHandler(
	// u domain.UserService,
	p domain.ProductService,
	// n domain.NotificationService,
	r domain.ReviewService,
) domain.UserViewHandler {
	return &userViewHandler{
		// userService:         u,
		productService: p,
		// notificationService: n,
		reviewService: r,
	}
}

func (u *userViewHandler) PageLanding(c echo.Context) error {
	data := pongo2.Context{}
	statusCode := http.StatusOK
	// get page
	pageString := c.QueryParam("page")
	page, _ := strconv.Atoi(pageString)
	if page <= 0 {
		page = 1
	}
	// get user data login
	username := c.Get("username")
	profilePic := c.Get("profile_pic")
	if username != nil {
		data["username"] = username
	}
	if profilePic != nil {
		data["profile"] = profilePic
	}
	// calling service
	reviews, pagination, err := u.reviewService.GetAllReviews(c.Request().Context(), page)
	if err != nil {
		statusCode = http.StatusInternalServerError
		data["status_code"] = statusCode
		data["err"] = err.Error()
		data["message"] = "Internal Server Error"

		return c.Render(http.StatusInternalServerError, config.FromViews("/error.html"), data)
	}
	// send html with data
	statusCode = http.StatusOK
	data["reviews"] = reviews
	data["user"] = c.Get("user")
	data["pagination"] = pagination

	return c.Render(statusCode, config.FromViews("/landing.html"), data)
}
func (u *userViewHandler) PageLogin(c echo.Context) error {
	data := pongo2.Context{}
	var credError entity.CredentialsError
	statusCode := http.StatusOK

	user_error, _ := c.Cookie("user_error")
	if user_error != nil {
		credError.UsernameError = user_error.Value
		user_error.Expires = time.Unix(0, 0)
		user_error.MaxAge = -1
		c.SetCookie(
			user_error,
		)
	}
	pass_error, _ := c.Cookie("pass_error")
	if pass_error != nil {
		credError.PasswordError = pass_error.Value
		pass_error.Expires = time.Unix(0, 0)
		pass_error.MaxAge = -1
		c.SetCookie(
			pass_error,
		)
	}

	data["error"] = credError

	return c.Render(statusCode, config.FromViews("/login.html"), data)
}
func (u *userViewHandler) PageRegister(c echo.Context) error {
	data := pongo2.Context{}
	var registerCredError entity.RegisterCredError
	statusCode := http.StatusOK

	// checking cookie error data if there any
	user_error, _ := c.Cookie("user_error")
	if user_error != nil {
		registerCredError.UsernameError = user_error.Value
		user_error.Expires = time.Unix(0, 0)
		user_error.MaxAge = -1
		c.SetCookie(
			user_error,
		)
	}
	pass_error, _ := c.Cookie("pass_error")
	if pass_error != nil {
		registerCredError.PasswordError = pass_error.Value
		pass_error.Expires = time.Unix(0, 0)
		pass_error.MaxAge = -1
		c.SetCookie(
			pass_error,
		)
	}
	confirm_pass_error, _ := c.Cookie("confirm_pass_error")
	if confirm_pass_error != nil {
		registerCredError.ConfirmPasswordError = confirm_pass_error.Value
		confirm_pass_error.Expires = time.Unix(0, 0)
		confirm_pass_error.MaxAge = -1
		c.SetCookie(
			confirm_pass_error,
		)
	}

	data["error"] = registerCredError
	return c.Render(statusCode, config.FromViews("/register.html"), data)
}
func (u *userViewHandler) PageProductSearch(c echo.Context) error {
	var data pongo2.Context
	statusCode := http.StatusOK
	return c.Render(statusCode, config.FromViews("/register.html"), data)
}
func (u *userViewHandler) PageUserProfile(c echo.Context) error {
	var data pongo2.Context
	statusCode := http.StatusOK
	return c.Render(statusCode, config.FromViews("/register.html"), data)
}
