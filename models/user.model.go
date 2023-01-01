package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string   `json:"username,omitempty" gorm:"unique"`
	Password      string   `json:"password,omitempty"`
	Reviews       []Review `json:"reviews,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FollowerTotal int      `json:"follower_total,omitempty"`
	ReviewTotal   int      `json:"review_total,omitempty"`
	Bio           string   `json:"bio,omitempty"`
}
