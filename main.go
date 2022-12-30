package main

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
)

// func init() {

// }

func main() {
	var tplExample = utility.Renderer{Debug: true}
	// set echo mux
	mux := echo.New()
	mux.Static("/assets", "public")
	mux.GET("/", func(c echo.Context) error {
		statusCode := http.StatusOK
		data := pongo2.Context{
			"user":     c.Get("user"),
			"main_url": "localhost:1323",
			"reviews": []entity.Review{
				{
					User: entity.User{
						Username: "Henzo",
					},
				},
			},
		}
		return c.Render(statusCode, config.FromViews("/landing.html"), data)
	})
	mux.Renderer = tplExample
	mux.Logger.Fatal(mux.Start(":1323"))
}
