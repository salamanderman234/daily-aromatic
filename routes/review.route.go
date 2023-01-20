package route

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type reviewRoute struct {
	mustAuthGroup *echo.Group
	handler       domain.ReviewHandler
}

func NewReviewRoute(mag *echo.Group, h domain.ReviewHandler) domain.Route {
	return &reviewRoute{
		mustAuthGroup: mag,
		handler:       h,
	}
}

func (r *reviewRoute) Register() {
	r.mustAuthGroup.POST("review", r.handler.CreateReviewProcess)
}
