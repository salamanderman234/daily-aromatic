package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID        int    `json:"user_id"`
	Username      string `json:"username"`
	FollowerTotal int    `json:"follower_total"`
	ReviewTotal   int    `json:"review_total"`
	Bio           string `json:"bio"`
}
