package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ProductID uint    `json:"-"`
	UserID    uint    `json:"-"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Rate      float64 `json:"rate"`
	Comment   string  `json:"comment"`
}
