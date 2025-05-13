package repository

import (
	"golang-gin/entity"

	"gorm.io/gorm"
)

type StudentRepository interface {
	InsertStudent(req *entity.Student) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) InsertStudent(student *entity.Student) error {
	err := r.db.Create(&student).Error

	return err
}
