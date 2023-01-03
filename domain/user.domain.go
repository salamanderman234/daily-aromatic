package domain

import (
	"context"

	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type UserService interface {
	CreateUser(c context.Context, user entity.User) error
	UpdateUser(c context.Context, id uint, updatedField entity.User) error
	GetUser(c context.Context, username string, pass string) (entity.User, bool, error)
	GetUserReviews(c context.Context, id uint) ([]entity.Review, error)
}

type UserRepository interface {
	CreateUser(c context.Context, user model.User) error
	UpdateUser(c context.Context, id uint, updatedField model.User) error
	GetUserByID(c context.Context, id uint) (model.User, error)
	GetUserByCred(c context.Context, username string, pass string) (model.User, bool, error)
}
