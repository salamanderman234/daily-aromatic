package entity

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/salamanderman234/daily-aromatic/constanta"
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
	if c.Username == "" {
		credError.UsernameError = constanta.EmptyCreds.Error()
		credError.ErrorCookies = append(credError.ErrorCookies, &http.Cookie{
			Name:  "user_error",
			Value: constanta.EmptyCreds.Error(),
			Path:  "/",
		})
	}
	if c.Password == "" {
		credError.PasswordError = constanta.EmptyCreds.Error()
		credError.ErrorCookies = append(credError.ErrorCookies, &http.Cookie{
			Name:  "pass_error",
			Value: constanta.EmptyCreds.Error(),
			Path:  "/",
		})
	}

	if credError.UsernameError != "" || credError.PasswordError != "" {
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
	if c.Username == "" {
		credError.ErrorCookies = append(credError.ErrorCookies, &http.Cookie{
			Name:  "user_error",
			Value: constanta.EmptyCreds.Error(),
			Path:  "/",
		})
	}

	if c.Password == "" {
		credError.ErrorCookies = append(credError.ErrorCookies, &http.Cookie{
			Name:  "pass_error",
			Value: constanta.EmptyCreds.Error(),
			Path:  "/",
		})
	}
	if c.ConfirmPassword == "" {
		credError.ErrorCookies = append(credError.ErrorCookies, &http.Cookie{
			Name:  "confirm_pass_error",
			Value: constanta.EmptyCreds.Error(),
			Path:  "/",
		})
	}

	if (c.ConfirmPassword != "" && c.Password != "") && (c.ConfirmPassword != c.Password) {
		credError.ErrorCookies = append(credError.ErrorCookies, &http.Cookie{
			Name:  "pass_error",
			Value: constanta.PasswordNotMatch.Error(),
			Path:  "/",
		})
		credError.ErrorCookies = append(credError.ErrorCookies, &http.Cookie{
			Name:  "confirm_pass_error",
			Value: constanta.PasswordNotMatch.Error(),
			Path:  "/",
		})
	}

	if len(credError.ErrorCookies) > 0 {
		return false, credError
	}
	return true, credError
}
