package entity

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

type JWTClaims struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	ProfilePic string `json:"profile_pic"`
	jwt.RegisteredClaims
}

// credentials
type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
type CredentialsError struct {
	UsernameError string
	PasswordError string
	ErrorCookies  []*http.Cookie
}

func (c *Credentials) Check() (bool, CredentialsError) {
	credError := CredentialsError{}
	variable.EmptyFieldValidator(c.Username, variable.UsernameErrCookie, &credError.ErrorCookies)
	variable.EmptyFieldValidator(c.Password, variable.PasswordErrCookie, &credError.ErrorCookies)
	if len(credError.ErrorCookies) > 0 {
		return false, credError
	}
	return true, credError
}

// register cred entity
type RegisterCred struct {
	Username        string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}
type RegisterCredError struct {
	UsernameError        string
	PasswordError        string
	ConfirmPasswordError string
	ErrorCookies         []*http.Cookie
}

func (c *RegisterCred) Check() (bool, RegisterCredError) {
	credError := RegisterCredError{}

	variable.EmptyFieldValidator(c.Username, variable.UsernameErrCookie, &credError.ErrorCookies)
	variable.MustMatchAndNotEmptyValidator(c.Password, c.ConfirmPassword, variable.PasswordErrCookie, variable.ConfirmPassErrCookie, &credError.ErrorCookies)
	if len(credError.ErrorCookies) > 0 {
		return false, credError
	}
	return true, credError
}
