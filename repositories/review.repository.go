package repository

import (
	"context"
	"errors"

	"github.com/salamanderman234/daily-aromatic/domain"
	model "github.com/salamanderman234/daily-aromatic/models"
	variable "github.com/salamanderman234/daily-aromatic/vars"
	"gorm.io/gorm"
)

type reviewRepository struct {
	conn  *gorm.DB
	table string
}

func NewReviewRepostory(c *gorm.DB) domain.ReviewRepository {
	return &reviewRepository{
		conn:  c,
		table: "reviews",
	}
}

func (r *reviewRepository) IsUserReviewLegal(c context.Context, userId uint) (bool, error) {
	result := r.conn.WithContext(c).Where("user_id = ?", userId).First(&model.Review{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, variable.ErrDataNotFound
		}
		return false, result.Error
	}
	return true, nil
}

func (r *reviewRepository) CreateReviews(c context.Context, reviews []model.Review) error {
	// create reviews and update rate to product
	r.conn.Transaction(func(tx *gorm.DB) error {
		product := model.Product{}
		for _, review := range reviews {
			result := r.conn.WithContext(c).Where("id = ?", review.ProductID).First(&product)
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}
			result = r.conn.WithContext(c).Create(&review)
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}
			product.JumlahReview++
			product.TotalRate += review.Rate
			product.Rating = product.TotalRate / float64(product.JumlahReview)
			r.conn.Save(&product)
		}
		tx.Commit()
		return nil
	})
	return nil
}

func (r *reviewRepository) GetTotalReview(c context.Context, filter model.Review) int64 {
	result := int64(0)
	var review model.Review
	r.conn.WithContext(c).Model(&review).Where(&filter).Count(&result)
	return result
}

func (r *reviewRepository) GetReviewByID(c context.Context, id uint) (model.Review, error) {
	var review model.Review
	result := r.conn.WithContext(c).Joins("Product").Joins("User").Where("id = ?", id).First(&review)
	if result.Error != nil {
		return model.Review{}, result.Error
	}
	return review, nil
}

func (r *reviewRepository) GetReviews(c context.Context, limit int, skip int, filter model.Review) ([]model.Review, error) {
	var reviews []model.Review
	result := r.conn.WithContext(c).
		Joins("Product").
		Joins("User").
		Where(&filter).
		Offset(skip).
		Limit(limit).
		Order("created_at desc").
		Find(&reviews)
	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}

func (r *reviewRepository) UpdateReview(c context.Context, id uint, updatedField model.Review) error {
	review := model.Review{}
	result := r.conn.WithContext(c).Model(&review).Where("id = ?", id).Updates(updatedField)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *reviewRepository) DeleteReview(c context.Context, id uint) error {
	review := model.Review{}
	result := r.conn.WithContext(c).Where("id = ?", id).Delete(&review)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
