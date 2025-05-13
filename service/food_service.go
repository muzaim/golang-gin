package service

import (
	"golang-gin/dto"
	"golang-gin/entity"
	"golang-gin/errorhandler"
	"golang-gin/repository"
)

type FoodService interface {
	CreateFood(req *dto.FoodRequest) error
}

type foodService struct {
	repository repository.FoodRepository
}

func NewFoodService(r repository.FoodRepository) *foodService {
	return &foodService{
		repository: r,
	}
}

func (s *foodService) CreateFood(req *dto.FoodRequest) error {
	if foodExist := s.repository.FoodNameExist(req.Name); foodExist {
		return &errorhandler.BadRequestError{Message: "Food already registered"}
	}

	food := entity.Food{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
	}

	if err := s.repository.CreateFood(&food); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	return nil
}
