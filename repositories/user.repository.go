package repository

import (
	"github.com/salamanderman234/daily-aromatic/domain"
	model "github.com/salamanderman234/daily-aromatic/models"
	"gorm.io/gorm"
)

type userRepository struct {
	conn  *gorm.DB
	table string
}

func NewUserRepository(c *gorm.DB) domain.UserRepository {
	return &userRepository{
		conn:  c,
		table: "users",
	}
}

func (u *userRepository) CreateUser(user model.User) error {
	result := u.conn.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *userRepository) UpdateUser(id uint, updatedField model.User) error {
	user := model.User{}
	result := u.conn.Model(&user).Where("id = ?", id).Updates(updatedField)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *userRepository) GetUserByID(id uint) (model.User, error) {
	user := model.User{}
	result := u.conn.Where("id = ?").First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *userRepository) GetUserByIDWithReviews(id uint) (model.User, error) {
	user := model.User{}
	result := u.conn.Preload("Reviews").Where("id = ?").First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *userRepository) GetUserByCred(username string, pass string) (model.User, bool, error) {
	cred := model.User{
		Username: username,
		Password: pass,
	}
	user := model.User{}
	result := u.conn.Where(&cred).First(&user)
	if result.Error != nil {
		return user, false, result.Error
	}

	return user, true, nil
}
