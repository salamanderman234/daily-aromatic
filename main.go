package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
	handler "github.com/salamanderman234/daily-aromatic/handlers"
	repository "github.com/salamanderman234/daily-aromatic/repositories"
	route "github.com/salamanderman234/daily-aromatic/routes"
	service "github.com/salamanderman234/daily-aromatic/services"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
	"github.com/spf13/viper"
)

func init() {
	// set config
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	// connect db
	conn, err := config.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	// // connect redis
	// rdb := config.ConnectToRedis()
	// fmt.Println(rdb)

	// set templating engine
	tplExample := utility.Renderer{Debug: true}
	// set echo mux
	mux := echo.New()
	mux.Renderer = tplExample
	mux.Static("/assets", "public")
	// mux.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	// // dependecies inject
	// // repo
	productRepo := repository.NewProductRepository(conn)
	// // service
	productServ := service.NewProductService(productRepo)
	// handler
	userViewHandler := handler.NewUserViewHandler(productServ)
	// route
	var routeList []domain.Route
	routeList = append(routeList, route.NewUserViewRoute(mux, userViewHandler))

	for _, route := range routeList {
		route.Register()
	}

	mux.Logger.Fatal(mux.Start(":1323"))
}
