package entity

import "github.com/golang-jwt/jwt/v4"

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type CredentialsError struct {
	UsernameError string
	PasswordError string
}

type RegisterCred struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

type RegisterCredError struct {
	EmailError           string
	UsernameError        string
	PasswordError        string
	ConfirmPasswordError string
}

type JWTClaims struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	ProfilePic string `json:"profile_pic"`
	jwt.RegisteredClaims
}
