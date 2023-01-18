package utility

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

// generating standart error page (status code, generated data, error page file name)
func ErrorPageFactory(status int) (int, pongo2.Context, string) {
	errHtml := "/error.html"
	data := pongo2.Context{
		"status_code": status,
	}
	if status == http.StatusNotFound {
		data["err"] = variable.ErrNotFound.Error()
	} else if status == http.StatusInternalServerError {
		data["err"] = variable.ErrInternalServer.Error()
	} else if status == http.StatusBadRequest {
		data["err"] = variable.ErrBadRequest.Error()
	} else if status == http.StatusUnauthorized {
		data["err"] = variable.ErrUnauthorized.Error()
	}

	return status, data, errHtml
}

// generating user data pongo context from with token or must auth request
func UserDataFactory(c echo.Context, data pongo2.Context) {
	username := c.Get("username")
	profilePic := c.Get("profile_pic")
	id := c.Get("id")
	if id != nil {
		data["id"] = id.(uint)
	}
	if username != nil {
		data["username"] = username.(string)
	}
	if profilePic != nil {
		data["profile"] = profilePic.(string)
	}
}

// generating expired cookie from given cookie
func ExpiredCookieFactory(ctx echo.Context, c *http.Cookie) {
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	ctx.SetCookie(c)
}
