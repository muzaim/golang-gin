package service

import (
	"golang-gin/dto"
	"golang-gin/entity"
	"golang-gin/errorhandler"
	"golang-gin/helper"
	"golang-gin/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "Email already registered"}
	}

	if nameExist := s.repository.NameExist(req.Name); nameExist {
		return &errorhandler.BadRequestError{Message: "Name already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{Message: "Password not match"}
	}

	passwordHash, err := helper.HashPasword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{}
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
		Gender:   req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
