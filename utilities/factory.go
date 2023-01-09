package utility

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/constanta"
)

func ErrorPageFactory(status int) (int, pongo2.Context, string) {
	errHtml := "/error.html"
	data := pongo2.Context{
		"status_code": status,
	}
	if status == http.StatusNotFound {
		data["err"] = constanta.NotFoundError.Error()
	} else if status == http.StatusInternalServerError {
		data["err"] = constanta.InternalServerError.Error()
	} else if status == http.StatusBadRequest {
		data["err"] = constanta.BadRequestError.Error()
	}

	return status, data, errHtml
}

func UserDataFactory(c echo.Context, data pongo2.Context) {
	username := c.Get("username")
	profilePic := c.Get("profile_pic")
	if username != nil {
		data["username"] = username
	}
	if profilePic != nil {
		data["profile"] = profilePic
	}
}

func ExpiredCookieFactory(ctx echo.Context, c *http.Cookie) {
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	ctx.SetCookie(c)
}
