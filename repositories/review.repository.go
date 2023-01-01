package repository

import (
	"github.com/salamanderman234/daily-aromatic/domain"
	model "github.com/salamanderman234/daily-aromatic/models"
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

func (r *reviewRepository) CreateReviews(reviews []model.Review) error {
	result := r.conn.Create(&reviews)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *reviewRepository) GetReviewByID(id uint) (model.Review, error) {
	var review model.Review
	result := r.conn.Joins("Product").Joins("User").Where("id = ?", id).First(&review)
	if result.Error != nil {
		return model.Review{}, result.Error
	}
	return review, nil
}

func (r *reviewRepository) GetReviews(limit int, skip int, filter model.Review) ([]model.Review, error) {
	var reviews []model.Review
	result := r.conn.Joins("Product").Joins("User").Where(&filter).Offset(skip).Limit(limit).Order("age desc, name").Find(&reviews)
	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}

func (r *reviewRepository) UpdateReview(id uint, updatedField model.Review) error {
	review := model.Review{}
	result := r.conn.Model(&review).Where("id = ?", id).Updates(updatedField)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *reviewRepository) DeleteReview(id uint) error {
	review := model.Review{}
	result := r.conn.Where("id = ?", id).Delete(&review)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
