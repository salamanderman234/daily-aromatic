package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	User    User    `json:"user"`
	Rate    float64 `json:"rate"`
	Product Product `json:"product"`
	Comment string  `json:"comment"`
}
