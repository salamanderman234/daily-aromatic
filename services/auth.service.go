package service

import (
	"context"
	"fmt"

	"github.com/salamanderman234/daily-aromatic/domain"
	model "github.com/salamanderman234/daily-aromatic/models"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
	variable "github.com/salamanderman234/daily-aromatic/vars"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo domain.UserRepository
}

func NewAuthService(r domain.UserRepository) domain.AuthService {
	return &authService{
		repo: r,
	}
}

func (a *authService) Login(c context.Context, username string, password string) (string, error) {
	user, err := a.repo.GetuserByUsername(c, username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", variable.ErrUserCredsNotFound
	}
	return utility.CreateJWT(user)

}
func (a *authService) Register(c context.Context, username string, password string) (string, error) {
	user := model.User{
		Username: username,
		Password: password,
		Avatar:   fmt.Sprintf("https://api.dicebear.com/5.x/avataaars-neutral/png?seed=%s&backgroundColor=ff8b3d&backgroundType=gradientLinear,solid&radius=50", username),
	}
	resultUser, err := a.repo.CreateUser(c, user)
	if err != nil {
		return "", err
	}

	return utility.CreateJWT(resultUser)
}

func (a *authService) IsUsernameAlreadyExists(c context.Context, username string) error {
	user, err := a.repo.GetuserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != "" {
		return variable.ErrMustUniqueUsername
	}
	return nil
}
