package service

import (
	"bootcamp/data"
	"bootcamp/model"
)

type IUsersService interface {
	Users() ([]model.Users, error)
	User(username string) (model.Users, error)
	NewUser(model.Users) error
	UpdateUser(user model.Users) error
}

type UsersService struct {
	Data data.IUsersData
}

func NewService(data data.IUsersData) IUsersService {
	return &UsersService{Data: data}
}

func (s *UsersService) Users() ([]model.Users, error) {
	users, err := s.Data.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UsersService) User(username string) (model.Users, error) {
	user, err := s.Data.GetByUsername(username)
	if err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func (s *UsersService) NewUser(user model.Users) error {
	return s.Data.Insert(user)
}

func (s *UsersService) UpdateUser(user model.Users) error {
	return s.Data.Update(user)
}
