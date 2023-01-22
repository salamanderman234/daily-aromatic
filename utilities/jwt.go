package utility

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	model "github.com/salamanderman234/daily-aromatic/models"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

type JWTClaims struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	ProfilePic string `json:"profile_pic"`
	jwt.RegisteredClaims
}

const (
	validityRange = 24
)

func CreateJWT(user model.User) (string, error) {
	claims := JWTClaims{
		ID:         user.ID,
		Username:   user.Username,
		ProfilePic: user.Avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(validityRange * time.Hour)},
		},
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwt.SignedString([]byte(config.GetJWTSecret()))
}

func JWTValidation(token string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecret()), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return claims, variable.ErrTokenNotValid
		}
		return claims, err
	}

	if !tkn.Valid {
		return claims, variable.ErrTokenExpired
	}

	return claims, nil
}
