package main

import (
	"context"
	"fmt"

	"github.com/salamanderman234/daily-aromatic/config"
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
	rdb := config.ConnectToRedis()
	val, _ := rdb.Get(context.Background(), "pepel").Result()
	fmt.Println(val)
	fmt.Println(rdb)
}
