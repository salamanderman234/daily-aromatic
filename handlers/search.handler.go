package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/daily-aromatic/domain"
)

type searchHandler struct {
	serv domain.ProductService
}

func NewSearchHandler(s domain.ProductService) domain.SearchHandler {
	return &searchHandler{
		serv: s,
	}
}

func (s *searchHandler) ProductSearchProcess(c echo.Context) error {
	return nil
}
