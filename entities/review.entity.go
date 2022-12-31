package entity

import "time"

type Review struct {
	User      User      `json:"user"`
	Rate      float64   `json:"rate"`
	Product   Product   `json:"product"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
