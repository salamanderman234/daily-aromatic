package main

import (
	"fmt"

	"github.com/salamanderman234/daily-aromatic/config"
	model "github.com/salamanderman234/daily-aromatic/models"
	"github.com/spf13/viper"
)

func init() {
	// set config
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	conn, err := config.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	modelList := []interface{}{
		&model.Product{},
		&model.User{},
		&model.Review{},
	}

	for _, model := range modelList {
		err = conn.AutoMigrate(model)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("migration success!")
	s := Seeder{}
	err = s.Generate(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("seed success!")
}
