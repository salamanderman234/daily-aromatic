package service

import (
	"context"
	"encoding/json"

	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) domain.UserService {
	return &userService{
		repo: r,
	}
}

func (u *userService) CreateUser(c context.Context, user entity.User) error {
	var userModel model.User
	temp, _ := json.Marshal(user)
	json.Unmarshal(temp, &userModel)

	_, err := u.repo.CreateUser(c, userModel)
	if err != nil {
		return err
	}

	return nil
}
func (u *userService) UpdateUser(c context.Context, id uint, updatedField entity.User) (string, error) {
	var updatedFieldModel model.User
	temp, _ := json.Marshal(updatedField)
	json.Unmarshal(temp, &updatedFieldModel)

	err := u.repo.UpdateUser(c, id, updatedFieldModel)
	if err != nil {
		return "", err
	}
	user, _ := u.repo.GetUserByID(c, id)

	return utility.CreateJWT(user)
}
func (u *userService) GetUser(c context.Context, username string) (entity.User, bool, error) {
	user, err := u.repo.GetuserByUsername(c, username)
	if err != nil {
		return entity.User{}, false, err
	}

	var userEntity entity.User
	temp, _ := json.Marshal(user)
	json.Unmarshal(temp, &userEntity)

	return userEntity, true, nil
}
func (a *userService) IsUsernameAlreadyExists(c context.Context, username string) error {
	user, err := a.repo.GetuserByUsername(c, username)
	if err != nil {
		return err
	}

	if user.Username != "" {
		return variable.ErrMustUniqueUsername
	}
	return nil
}
