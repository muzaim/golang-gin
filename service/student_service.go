package service

import (
	"golang-gin/dto"
	"golang-gin/entity"
	"golang-gin/errorhandler"
	"golang-gin/repository"
)

type StudentService interface {
	InsertStudent(req *dto.StudentRequest) error
}

type studentService struct {
	repository repository.StudentRepository
}

func NewStudentService(r repository.StudentRepository) *studentService {
	return &studentService{
		repository: r,
	}
}

func (s *studentService) InsertStudent(req *dto.StudentRequest) error {
	student := entity.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Major:     req.Major,
		Email:     req.Email,
	}

	if err := s.repository.InsertStudent(&student); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
