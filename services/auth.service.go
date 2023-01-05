package service

import (
	"context"

	"github.com/salamanderman234/daily-aromatic/constanta"
	"github.com/salamanderman234/daily-aromatic/domain"
	model "github.com/salamanderman234/daily-aromatic/models"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return "", constanta.UserDataNotFoundWithCreds
	}
	return utility.CreateJWT(user)

}
func (a *authService) Register(c context.Context, username string, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := model.User{
		Username: username,
		Password: string(hashedPassword),
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
		return constanta.UsernameAlreadyExists
	}
	return nil
}
