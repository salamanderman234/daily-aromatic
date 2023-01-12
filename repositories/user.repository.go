package repository

import (
	"context"
	"errors"

	"github.com/salamanderman234/daily-aromatic/domain"
	model "github.com/salamanderman234/daily-aromatic/models"
	"golang.org/x/crypto/bcrypt"
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

func (u *userRepository) CreateUser(c context.Context, user model.User) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hashedPassword)
	result := u.conn.WithContext(c).Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
func (u *userRepository) UpdateUser(c context.Context, id uint, updatedField model.User) error {
	user := model.User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedField.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	updatedField.Password = string(hashedPassword)
	result := u.conn.WithContext(c).Model(&user).Where("id = ?", id).Updates(updatedField)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (u *userRepository) GetUserByID(c context.Context, id uint) (model.User, error) {
	user := model.User{}
	result := u.conn.WithContext(c).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *userRepository) GetUserByIDWithReviews(c context.Context, id uint) (model.User, error) {
	user := model.User{}
	result := u.conn.WithContext(c).Preload("Reviews").Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *userRepository) GetUserByCred(c context.Context, username string, pass string) (model.User, bool, error) {
	cred := model.User{
		Username: username,
		Password: pass,
	}
	user := model.User{}
	result := u.conn.WithContext(c).Where(&cred).First(&user)
	if result.Error != nil {
		return user, false, result.Error
	}

	return user, true, nil
}

func (u *userRepository) GetuserByUsername(c context.Context, username string) (model.User, error) {
	user := model.User{}
	// get user with reviews limit 10
	result := u.conn.WithContext(c).Preload("Reviews", func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(10)
	}).Preload("Reviews.Product").Where("username = ?", username).First(&user)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, result.Error
		}
	}
	return user, nil
}
