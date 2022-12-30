package domain

import (
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type ReviewService interface {
	CreateReview(review entity.Review) error
	GetAllReviews() ([]entity.Review, error)
	GetUserReviews(user_id uint) ([]entity.Review, error)
	GetReview(id uint) (entity.Review, error)
	UpdateReview(id uint, updatedField entity.Review) error
	DeleteReview(id uint) error
}

type ReviewRepository interface {
	CreateReview(review model.Review) error
	GetReviewByID(id uint) (model.Review, error)
	GetAllReviews() ([]model.Review, error)
	GetReviewsByUserID(user_id uint) ([]model.Review, error)
	UpdateReview(id uint) error
	DeleteReview(id uint) error
}
