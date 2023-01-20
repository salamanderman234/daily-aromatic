package entity

import (
	"net/http"

	variable "github.com/salamanderman234/daily-aromatic/vars"
)

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

func (c *Credentials) Validate() (bool, CredentialsError) {
	credError := CredentialsError{}
	emptyFieldValidator(c.Username, variable.UsernameErrCookie, &credError.ErrorCookies)
	emptyFieldValidator(c.Password, variable.PasswordErrCookie, &credError.ErrorCookies)
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

func (c *RegisterCred) Validate() (bool, RegisterCredError) {
	credError := RegisterCredError{}

	emptyFieldValidator(c.Username, variable.UsernameErrCookie, &credError.ErrorCookies)
	// variable.PasswordValidator(c.Password, variable.PasswordErrCookie, &credError.ErrorCookies)
	mustMatchAndNotEmptyValidator(c.Password, c.ConfirmPassword, variable.PasswordErrCookie, variable.ConfirmPassErrCookie, &credError.ErrorCookies)
	if len(credError.ErrorCookies) > 0 {
		return false, credError
	}
	return true, credError
}
