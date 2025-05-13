package repository

import (
	"golang-gin/entity"
	"log"

	"gorm.io/gorm"
)

type FoodRepository interface {
	FoodNameExist(name string) bool
	CreateFood(req *entity.Food) error
}

type foodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *foodRepository {
	return &foodRepository{
		db: db,
	}
}

func (r *foodRepository) CreateFood(food *entity.Food) error {
	err := r.db.Create(&food).Error

	return err
}

func (r *foodRepository) FoodNameExist(name string) bool {
	var food entity.Food
	err := r.db.First(&food, "name = ?", name).Error

	if err != nil {
		log.Printf("[NOT EXIST] '%s' not found, err: %v", name, err)
		return false
	}

	log.Printf("[EXIST] Found food: %+v", food)
	return true
}
