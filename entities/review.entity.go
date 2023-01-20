package entity

import "time"

type Review struct {
	User      User      `json:"user"`
	Rate      float64   `json:"rate"`
	Product   Product   `json:"product"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewForm struct {
	UserID    uint    `json:"user_id" form:"user_id"`
	ProductID uint    `json:"product_id" form:"product_id"`
	Rate      float64 `json:"rate" form:"rate"`
	Comment   string  `json:"comment" form:"comment"`
}
