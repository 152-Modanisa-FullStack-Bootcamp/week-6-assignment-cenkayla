package data

import (
	"bootcamp/model"
	"errors"
)

type IUsersData interface {
	Insert(model.Users) error
	Update(model.Users) error
	GetByUsername(username string) (model.Users, error)
	GetAll() ([]model.Users, error)
}

type UsersData struct {
	database map[string]int
}

func NewInMemoryData() IUsersData {
	return &UsersData{database: make(map[string]int)}
}

func (u *UsersData) Insert(user model.Users) error {
	u.database[user.Username] = user.Balance

	if user.Username == "" {
		return errors.New("Could not insert user")
	}

	return nil
}

func (u *UsersData) Update(user model.Users) error {
	u.database[user.Username] = user.Balance

	if user.Username == "" {
		return errors.New("User not found")
	}
	return nil
}

func (u *UsersData) GetAll() ([]model.Users, error) {
	var users []model.Users

	for k, v := range u.database {
		users = append(users, model.Users{Username: k, Balance: v})
	}

	if len(users) == 0 {
		return nil, errors.New("No users found")
	}

	return users, nil

}

func (u *UsersData) GetByUsername(username string) (model.Users, error) {
	var user model.Users

	for k, v := range u.database {
		if k == username {
			user = model.Users{Username: k, Balance: v}
		}
	}

	if user.Username == "" {
		return model.Users{}, errors.New("User not found")
	}

	return user, nil
}
