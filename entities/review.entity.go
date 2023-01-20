package entity

import (
	"net/http"
	"time"

	variable "github.com/salamanderman234/daily-aromatic/vars"
)

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
type ReviewFormError struct {
	RateError    string
	CommentError string
	ErrorCookies []*http.Cookie
}

func (r *ReviewForm) Validate() (bool, ReviewFormError) {

	err := ReviewFormError{}
	// validating
	emptyFieldValidator(r.Comment, variable.CommentErrCookie, &err.ErrorCookies)
	mustEqualOrGreaterValidator(float64(len(r.Comment)), 10.0, variable.RateErrCookie, &err.ErrorCookies)
	mustInRangeValidator(r.Rate, 0.0, 5.0, variable.RateErrCookie, &err.ErrorCookies)
	// checking if there is an error
	if len(err.ErrorCookies) > 0 {
		return false, err
	}
	return true, err
}
