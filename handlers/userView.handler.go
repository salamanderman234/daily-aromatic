package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

type userViewHandler struct {
	userService    domain.UserService
	productService domain.ProductService
	// notificationService domain.NotificationService
	reviewService domain.ReviewService
}

func NewUserViewHandler(
	u domain.UserService,
	p domain.ProductService,
	// n domain.NotificationService,
	r domain.ReviewService,
) domain.UserViewHandler {
	return &userViewHandler{
		userService:    u,
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
	// get the error cookie and set it to expired
	user_error, _ := c.Cookie(variable.UsernameErrCookie)
	if user_error != nil {
		credError.UsernameError = user_error.Value
		utility.ExpiredCookieFactory(c, user_error)
	}
	pass_error, _ := c.Cookie(variable.PasswordErrCookie)
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
	user_error, _ := c.Cookie(variable.UsernameErrCookie)
	if user_error != nil {
		registerCredError.UsernameError = user_error.Value
		utility.ExpiredCookieFactory(c, user_error)
	}
	pass_error, _ := c.Cookie(variable.PasswordErrCookie)
	if pass_error != nil {
		registerCredError.PasswordError = pass_error.Value
		utility.ExpiredCookieFactory(c, pass_error)
	}
	confirm_pass_error, _ := c.Cookie(variable.ConfirmPassErrCookie)
	if confirm_pass_error != nil {
		registerCredError.ConfirmPasswordError = confirm_pass_error.Value
		utility.ExpiredCookieFactory(c, confirm_pass_error)
	}

	data["error"] = registerCredError
	return c.Render(statusCode, config.FromViews("/register.html"), data)
}
func (u *userViewHandler) PageProductSearch(c echo.Context) error {
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
	return c.Render(http.StatusOK, config.FromViews("/product-search.html"), data)
}

func (u *userViewHandler) PageUserProfile(c echo.Context) error {
	data := pongo2.Context{}
	utility.UserDataFactory(c, data)
	_, ok := data["username"]
	if !ok {
		status, data, fileName := utility.ErrorPageFactory(http.StatusUnauthorized)
		return c.Render(status, config.FromViews(fileName), data)
	}
	user, _, err := u.userService.GetUser(c.Request().Context(), data["username"].(string))
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}

	// checking cookie error data if there any
	var registerCredError entity.UserError
	user_error, _ := c.Cookie(variable.UsernameErrCookie)
	if user_error != nil {
		registerCredError.UsernameError = user_error.Value
		utility.ExpiredCookieFactory(c, user_error)
	}
	pass_error, _ := c.Cookie(variable.PasswordErrCookie)
	if pass_error != nil {
		registerCredError.PasswordError = pass_error.Value
		utility.ExpiredCookieFactory(c, pass_error)
	}
	data["error"] = registerCredError
	data["user"] = user
	return c.Render(http.StatusOK, config.FromViews("/profile.html"), data)
}

func (u *userViewHandler) PageDiffUserProfile(c echo.Context) error {
	data := pongo2.Context{}
	utility.UserDataFactory(c, data)
	username := c.Param("username")
	if username == "" {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
	}
	if username != "" && data["username"] != "" {
		if username == data["username"] {
			return c.Redirect(http.StatusFound, "/profile")
		}
	}
	user, _, err := u.userService.GetUser(c.Request().Context(), username)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}

	data["user"] = user
	return c.Render(http.StatusOK, config.FromViews("/diff-user-profile.html"), data)
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
		if errors.Is(err, variable.ErrDataNotFound) {
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

func (u *userViewHandler) NewReviewPage(c echo.Context) error {
	data := pongo2.Context{}
	utility.UserDataFactory(c, data)
	page := 1
	keyword := c.QueryParam("keyword")
	page, _ = strconv.Atoi(c.QueryParam("page"))
	productID, _ := strconv.Atoi(c.QueryParam("product"))
	// if product param is not 0 then show the create new review page
	if productID != 0 {
		product, err := u.productService.GetProduct(c.Request().Context(), uint(productID))
		if err != nil {
			if errors.Is(err, variable.ErrDataNotFound) {
				status, data, fileName := utility.ErrorPageFactory(http.StatusNotFound)
				return c.Render(status, config.FromViews(fileName), data)
			}
			status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
			return c.Render(status, config.FromViews(fileName), data)
		}
		data["product"] = product
		return c.Render(http.StatusOK, config.FromViews("/new-review.html"), data)
	}
	// otherwise show select product for review page
	// makin filter
	filter := entity.Product{
		Nama:     keyword,
		Pabrikan: keyword,
		Aroma:    keyword,
	}
	// calling serv
	id, ok := data["id"]
	if !ok {
		status, data, fileName := utility.ErrorPageFactory(http.StatusUnauthorized)
		return c.Render(status, config.FromViews(fileName), data)
	}
	result, pagination, err := u.productService.GetProductNotInUserReview(c.Request().Context(), page, filter, id.(uint))
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}

	// assigning data to view
	data["products"] = result
	// get user data login
	data["pagination"] = pagination
	data["keyword"] = keyword
	return c.Render(http.StatusOK, config.FromViews("/pilih-product.html"), data)
}
