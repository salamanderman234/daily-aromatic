package utility

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/constanta"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

const (
	validityRange = 24
)

func CreateJWT(user model.User) (string, error) {
	claims := entity.JWTClaims{
		ID:         user.ID,
		Username:   user.Username,
		ProfilePic: "test",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(validityRange * time.Hour)},
		},
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwt.SignedString([]byte(config.GetJWTSecret()))
}

func JWTValidation(token string) (*entity.JWTClaims, error) {
	claims := entity.JWTClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecret()), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return &claims, constanta.TokenInvalid
		}
		return &claims, err
	}

	if !tkn.Valid {
		return &claims, constanta.TokenExpired
	}

	return &claims, nil
}
