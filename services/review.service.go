package service

import (
	"context"
	"encoding/json"

	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
)

type reviewService struct {
	repo domain.ReviewRepository
}

func NewReviewService(r domain.ReviewRepository) domain.ReviewService {
	return &reviewService{
		repo: r,
	}
}

func (r *reviewService) CreateReview(c context.Context, review entity.ReviewForm) error {
	modelReview := model.Review{}
	temp, _ := json.Marshal(review)
	json.Unmarshal(temp, &modelReview)

	err := r.repo.CreateReviews(
		c,
		[]model.Review{modelReview},
	)
	if err != nil {
		return err
	}

	return nil
}
func (r *reviewService) GetAllReviews(c context.Context) ([]entity.Review, error) {
	result, err := r.repo.GetReviews()

}
func (r *reviewService) GetReview(c context.Context, id uint) (entity.Review, error) {

}
func (r *reviewService) UpdateReview(c context.Context, id uint, updatedField entity.Review) error {

}
func (r *reviewService) DeleteReview(c context.Context, id uint) error {

}
