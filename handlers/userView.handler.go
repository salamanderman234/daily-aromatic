package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/constanta"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
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
	// get page
	pageString := c.QueryParam("page")
	page, _ := strconv.Atoi(pageString)
	if page <= 0 {
		page = 1
	}
	// calling service
	reviews, pagination, err := u.reviewService.GetAllReviews(c.Request().Context(), page)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// send html with data
	data := pongo2.Context{}
	data["reviews"] = reviews
	data["pagination"] = pagination
	// get user data login
	utility.UserDataFactory(c, data)

	return c.Render(http.StatusOK, config.FromViews("/landing.html"), data)
}

func (u *userViewHandler) PageLogin(c echo.Context) error {
	data := pongo2.Context{}
	var credError entity.CredentialsError
	statusCode := http.StatusOK

	user_error, _ := c.Cookie("user_error")
	if user_error != nil {
		credError.UsernameError = user_error.Value
		utility.ExpiredCookieFactory(c, user_error)
	}
	pass_error, _ := c.Cookie("pass_error")
	if pass_error != nil {
		credError.PasswordError = pass_error.Value
		utility.ExpiredCookieFactory(c, pass_error)
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
		utility.ExpiredCookieFactory(c, user_error)
	}
	pass_error, _ := c.Cookie("pass_error")
	if pass_error != nil {
		registerCredError.PasswordError = pass_error.Value
		utility.ExpiredCookieFactory(c, pass_error)
	}
	confirm_pass_error, _ := c.Cookie("confirm_pass_error")
	if confirm_pass_error != nil {
		registerCredError.ConfirmPasswordError = confirm_pass_error.Value
		utility.ExpiredCookieFactory(c, confirm_pass_error)
	}

	data["error"] = registerCredError
	return c.Render(statusCode, config.FromViews("/register.html"), data)
}
func (u *userViewHandler) PageProductSearch(c echo.Context) error {
	statusCode := http.StatusOK
	page := 1
	// get params
	keyword := c.QueryParam("keyword")
	page, _ = strconv.Atoi(c.QueryParam("page"))
	// makin filter
	filter := entity.Product{
		Nama:     keyword,
		Pabrikan: keyword,
		Aroma:    keyword,
	}
	// calling serv
	result, pagination, err := u.productService.GetProductByFilter(c.Request().Context(), page, filter)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// assigning data to view
	data := pongo2.Context{}
	data["products"] = result
	// get user data login
	utility.UserDataFactory(c, data)
	data["pagination"] = pagination
	data["keyword"] = keyword
	return c.Render(statusCode, config.FromViews("/product-search.html"), data)
}

func (u *userViewHandler) PageUserProfile(c echo.Context) error {
	var data pongo2.Context
	statusCode := http.StatusOK
	return c.Render(statusCode, config.FromViews("/register.html"), data)
}

func (p *userViewHandler) ProductDetailPage(c echo.Context) error {
	id := c.Param("id")
	idInt := 0
	// checking if param is emtpy
	if id == "" {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// check if param is a valid id
	idInt, err := strconv.Atoi(id)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusNotFound)
		return c.Render(status, config.FromViews(fileName), data)
	}
	// calling service
	product, err := p.productService.GetProduct(c.Request().Context(), uint(idInt))
	if err != nil {
		if errors.Is(err, constanta.ProductNotFound) {
			status, data, fileName := utility.ErrorPageFactory(http.StatusNotFound)
			return c.Render(status, config.FromViews(fileName), data)
		} else {
			status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
			return c.Render(status, config.FromViews(fileName), data)
		}
	}
	data := pongo2.Context{}
	data["product"] = product
	// get user data login
	utility.UserDataFactory(c, data)
	return c.Render(http.StatusOK, config.FromViews("/product-detail.html"), data)
}
