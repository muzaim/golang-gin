package repository

import (
	"golang-gin/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
	Register(req *entity.User) error
	NameExist(name string) bool
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *entity.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *authRepository) EmailExist(email string) bool {
	var user entity.User

	err := r.db.First(&user, "email=?", &email).Error

	return err == nil
}

func (r *authRepository) NameExist(name string) bool {
	var user entity.User

	err := r.db.First(&user, "name=?", &name).Error

	return err == nil
}
