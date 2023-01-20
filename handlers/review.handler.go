package handler

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/config"
	"github.com/salamanderman234/daily-aromatic/domain"
	entity "github.com/salamanderman234/daily-aromatic/entities"
	utility "github.com/salamanderman234/daily-aromatic/utilities"
)

type reviewHandler struct {
	serv domain.ReviewService
}

func NewReviewHandler(s domain.ReviewService) domain.ReviewHandler {
	return &reviewHandler{
		serv: s,
	}
}

func (r *reviewHandler) CreateReviewProcess(c echo.Context) error {
	data := pongo2.Context{}
	newReview := entity.ReviewForm{}
	utility.UserDataFactory(c, data)
	if _, ok := data["id"]; !ok {
		status, data, fileName := utility.ErrorPageFactory(http.StatusUnauthorized)
		return c.Render(status, config.FromViews(fileName), data)
	}

	if err := c.Bind(&newReview); err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusBadRequest)
		return c.Render(status, config.FromViews(fileName), data)
	}
	newReview.UserID = data["id"].(uint)
	err := r.serv.CreateReview(c.Request().Context(), newReview)
	if err != nil {
		status, data, fileName := utility.ErrorPageFactory(http.StatusInternalServerError)
		return c.Render(status, config.FromViews(fileName), data)
	}
	return c.Redirect(http.StatusFound, "/")
}
