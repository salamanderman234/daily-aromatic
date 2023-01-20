package entity

import (
	"net/http"
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

func (u *User) Validate() (bool, UserError) {
	userErr := UserError{}
	emptyFieldValidator(u.Username, "user_error", &userErr.ErrorCookies)

	if len(userErr.ErrorCookies) > 0 {
		return false, userErr
	}
	return true, userErr
}
