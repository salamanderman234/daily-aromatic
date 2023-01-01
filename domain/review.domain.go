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
	CreateReviews(reviews []model.Review) error
	GetReviewByID(id uint) (model.Review, error)
	GetReviews(limit int, skip int, filter model.Review) ([]model.Review, error)
	UpdateReview(id uint, updatedField model.Review) error
	DeleteReview(id uint) error
}
