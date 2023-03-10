package service

import (
	"context"
	"encoding/json"
	"errors"
	"math"

	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	model "github.com/salamanderman234/daily-aromatic/models"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

type reviewService struct {
	repo        domain.ReviewRepository
	productRepo domain.ProductRepository
	userRepo    domain.UserRepository
}

func NewReviewService(r domain.ReviewRepository, p domain.ProductRepository, u domain.UserRepository) domain.ReviewService {
	return &reviewService{
		repo:        r,
		productRepo: p,
		userRepo:    u,
	}
}

func (r *reviewService) CreateReview(c context.Context, review entity.ReviewForm) error {
	// creating model
	modelReview := model.Review{}
	temp, _ := json.Marshal(review)
	json.Unmarshal(temp, &modelReview)

	// validation
	exists, err := r.productRepo.IsProductExists(c, review.ProductID)
	if !exists {
		return err
	}
	exists, err = r.userRepo.IsUserExists(c, review.UserID)
	if !exists {
		return err
	}
	legal, err := r.repo.IsUserReviewLegal(c, review.UserID)
	if !legal {
		return err
	}
	// calling repo
	err = r.repo.CreateReviews(
		c,
		[]model.Review{modelReview},
	)
	if err != nil {
		return err
	}

	return nil
}
func (r *reviewService) GetAllReviews(c context.Context, page int) ([]entity.Review, entity.Pagination, error) {
	var reviewsEntity []entity.Review
	limit := variable.QueryLimit
	offset := variable.DefaultOffset(page)
	// check if page exists
	maxPage := r.repo.GetTotalReview(c, model.Review{})
	pageMax := int(math.Ceil(float64(maxPage) / float64(limit)))
	if maxPage == 0 {
		maxPage = 1
	}
	if page > pageMax {
		return nil, entity.Pagination{}, errors.New("not found")
	}
	if page == 0 {
		page = 1
	}
	// calling repo
	reviewsModel, err := r.repo.GetReviews(c, limit, offset, model.Review{})
	if err != nil {
		return nil, entity.Pagination{}, err
	}
	// creating pagination obj
	pagination := entity.Pagination{
		CurrentPage:  page,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		MaxPage:      pageMax,
	}
	// to entity
	temp, _ := json.Marshal(reviewsModel)
	json.Unmarshal(temp, &reviewsEntity)
	return reviewsEntity, pagination, nil

}
func (r *reviewService) GetReview(c context.Context, id uint) (entity.Review, error) {
	var reviewEntity entity.Review
	// calling repo
	review, err := r.repo.GetReviewByID(c, id)
	if err != nil {
		return entity.Review{}, err
	}
	// to entity
	temp, _ := json.Marshal(review)
	json.Unmarshal(temp, &reviewEntity)

	return reviewEntity, nil
}
func (r *reviewService) UpdateReview(c context.Context, id uint, updatedField entity.Review) error {
	return nil
}
func (r *reviewService) DeleteReview(c context.Context, id uint) error {
	return nil
}
