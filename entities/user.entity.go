package entity

import (
	"net/http"

	variable "github.com/salamanderman234/daily-aromatic/vars"
)

type User struct {
	ID            uint           `json:"id"`
	Avatar        string         `json:"avatar,omitempty"`
	Username      string         `json:"username" form:"username"`
	Password      string         `form:"password"`
	FollowerTotal int            `json:"follower_total"`
	ReviewTotal   int            `json:"review_total"`
	Bio           string         `json:"bio" form:"bio"`
	Notifications []Notification `json:"notifications"`
	Reviews       []Review       `json:"reviews"`
}

type UserError struct {
	UsernameError string
	PasswordError string
	BioError      string
	ErrorCookies  []*http.Cookie
}

func (u *User) Check() (bool, UserError) {
	userErr := UserError{}
	variable.EmptyFieldValidator(u.Username, "user_error", &userErr.ErrorCookies)
	
	if len(userErr.ErrorCookies) > 0 {
		return false, userErr
	}
	return true, userErr
}
