package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	AppName = "Daily Aromatic"
)

func FromViews(path string) string {
	return fmt.Sprintf("./views%s", path)
}

func ConnectToDatabase() (*gorm.DB, error) {
	host := viper.GetString("database.host")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	port := viper.GetInt("database.port")
	name := viper.GetString("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, name)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ConnectToRedis() *redis.Client {
	user := viper.GetString("redis.username")
	passwd := viper.GetString("redis.password")
	host := viper.GetString("redis.host")
	port := viper.GetInt("redis.port")

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Username: user,
			Password: passwd,
		},
	)

	return rdb
}

func GetJWTSecret() string {
	return viper.GetString("jwt_secret")
}
