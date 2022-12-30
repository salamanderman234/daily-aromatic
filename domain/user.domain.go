package domain

import (
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type UserService interface {
	CreateUser(user entity.User) error
	UpdateUser(id uint, updatedField entity.User) error
	GetUser(username string, pass string) (entity.User, bool, error)
	GetUserReviews(id uint) ([]entity.Review, error)
}

type UserRepository interface {
	CreateUser(user model.User) error
	UpdateUser(id uint, updatedField model.User)
	GetUserByID(id uint) (model.User, error)
	GetUserByCred(username string, pass string) (entity.User, bool, error)
}
