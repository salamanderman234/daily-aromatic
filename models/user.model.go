package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Avatar        string   `json:"avatar,omitempty"`
	Username      string   `json:"username,omitempty" gorm:"unique"`
	Password      string   `json:"password,omitempty"`
	Reviews       []Review `json:"reviews,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FollowerTotal int      `json:"follower_total,omitempty"`
	ReviewTotal   int      `json:"review_total,omitempty"`
	Bio           string   `json:"bio,omitempty"`
}
