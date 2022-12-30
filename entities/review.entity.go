package entity

type Review struct {
	User    User    `json:"user"`
	Rate    float64 `json:"rate"`
	Product Product `json:"product"`
	Comment string  `json:"comment"`
}
