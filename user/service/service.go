package service

import (
	"crowdfund-go/entity"
	"crowdfund-go/user"
	"crowdfund-go/user/request"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository user.Repository
}

func NewService(repository user.Repository) *service {
	return &service{repository: repository}
}

func (s *service) Register(req request.Register) (entity.User, error) {
	user := entity.User{}
	user.Name = req.Name
	user.Email = req.Email
	user.Occupation = req.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(req request.Login) (entity.User, error) {
	email := req.Email
	password := req.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(req request.CheckEmail) (bool, error) {
	email := req.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (entity.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (entity.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that id")
	}

	return user, nil
}
