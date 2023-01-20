package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
	handler "github.com/salamanderman234/daily-aromatic/handlers"
	middleware "github.com/salamanderman234/daily-aromatic/middlewares"
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
	reviewRepo := repository.NewReviewRepostory(conn)
	userRepo := repository.NewUserRepository(conn)
	// // service
	productServ := service.NewProductService(productRepo)
	reviewServ := service.NewReviewService(reviewRepo)
	authServ := service.NewAuthService(userRepo)
	userServ := service.NewUserService(userRepo)
	// handler
	userViewHandler := handler.NewUserViewHandler(userServ, productServ, reviewServ)
	authHandler := handler.NewAuthHandler(authServ)
	userHandler := handler.NewUserHandler(userServ)
	reviewHandler := handler.NewReviewHandler(reviewServ)
	// route
	groupWithToken := mux.Group("/", middleware.WithToken)
	groupMustAuth := mux.Group("/", middleware.MustAuth)
	groupMustGuest := mux.Group("/", middleware.MustGuest)
	routeList := []domain.Route{
		route.NewUserViewRoute(groupWithToken, groupMustAuth, groupMustGuest, userViewHandler),
		route.NewAuthRoute(groupMustGuest, groupMustAuth, authHandler),
		route.NewUserRoute(groupMustAuth, userHandler),
		route.NewReviewRoute(groupMustAuth, reviewHandler),
	}

	for _, route := range routeList {
		route.Register()
	}

	mux.Logger.Fatal(mux.Start(":1323"))
}
