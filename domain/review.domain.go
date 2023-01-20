package domain

import (
	"context"

	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type ReviewService interface {
	CreateReview(c context.Context, review entity.ReviewForm) error
	GetAllReviews(c context.Context, page int) ([]entity.Review, entity.Pagination, error)
	GetReview(c context.Context, id uint) (entity.Review, error)
	UpdateReview(c context.Context, id uint, updatedField entity.Review) error
	DeleteReview(c context.Context, id uint) error
}

type ReviewRepository interface {
	IsUserReviewLegal(c context.Context, userId uint) (bool, error)
	GetTotalReview(c context.Context, filter model.Review) int64
	CreateReviews(c context.Context, reviews []model.Review) error
	GetReviewByID(c context.Context, id uint) (model.Review, error)
	GetReviews(c context.Context, limit int, skip int, filter model.Review) ([]model.Review, error)
	UpdateReview(c context.Context, id uint, updatedField model.Review) error
	DeleteReview(c context.Context, id uint) error
}
