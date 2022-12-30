package utility

import (
	"errors"
	"fmt"
	"io"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
)

type Renderer struct {
	templates *pongo2.Template
	ctx       pongo2.Context
	err       error
	Debug     bool
}

func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if data != nil {
		var ok bool
		r.ctx, ok = data.(pongo2.Context)
		if !ok {
			return errors.New("no pongo2.Context data was passed")
		}
	}

	if r.Debug {
		r.templates, r.err = pongo2.FromFile(name)
	} else {
		r.templates, r.err = pongo2.FromCache(name)
	}
	// Add some static values
	r.ctx["app_name"] = config.AppName
	if r.err != nil {
		fmt.Println(r.err)
		return r.err
	}
	return r.templates.ExecuteWriter(r.ctx, w)
}
