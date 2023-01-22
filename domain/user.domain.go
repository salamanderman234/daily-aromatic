package domain

import (
	"context"

	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type UserService interface {
	CreateUser(c context.Context, user entity.User) error
	UpdateUser(c context.Context, id uint, updatedField entity.User) (string, error)
	GetUser(c context.Context, username string) (entity.User, bool, error)
	IsUsernameAlreadyExists(c context.Context, username string) error
}

type UserRepository interface {
	IsUserExists(c context.Context, id uint) (bool, error)
	CreateUser(c context.Context, user model.User) (model.User, error)
	UpdateUser(c context.Context, id uint, updatedField model.User) error
	GetUserByID(c context.Context, id uint) (model.User, error)
	GetuserByUsername(c context.Context, username string) (model.User, error)
	GetUserByIDWithReviews(c context.Context, id uint) (model.User, error)
	GetUserByCred(c context.Context, username string, pass string) (model.User, bool, error)
}
