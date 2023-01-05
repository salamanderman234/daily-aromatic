package main

import (
	"fmt"

	"github.com/salamanderman234/daily-aromatic/config"
	model "github.com/salamanderman234/daily-aromatic/models"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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

func productSeed() []model.Product {
	productList := []model.Product{
		{
			Nama:     "Natsu Birla Kemenyan",
			ImageUrl: "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/6/19/f5c15304-a4e6-4334-b02f-4b3ee27c1501.jpg",
			Pabrikan: "Napura Dupa",
		},
	}
	return productList
}

func userSeed() []model.User {
	pass, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
	userList := []model.User{
		{
			Username: "salamanderman234",
			Password: string(pass),
		},
	}
	return userList
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

	conn.Create(userSeed())
	conn.Create(productSeed())
}
