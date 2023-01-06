package domain

import "github.com/labstack/echo/v4"

type UserViewHandler interface {
	PageLanding(c echo.Context) error
	PageLogin(c echo.Context) error
	PageRegister(c echo.Context) error
	PageProductSearch(c echo.Context) error
	PageUserProfile(c echo.Context) error
}

type AuthHandler interface {
	LoginProcess(c echo.Context) error
	RegisterProcess(c echo.Context) error
}

type SearchHandler interface {
	ProductSearchProcess(c echo.Context) error
}
