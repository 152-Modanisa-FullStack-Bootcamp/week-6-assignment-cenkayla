package service_test

import (
	"testing"

	"bootcamp/mock"
	"bootcamp/model"
	"bootcamp/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShouldDelegateUsers(t *testing.T) {
	mockData := mock.NewMockIUsersData(gomock.NewController(t))

	mockData.EXPECT().GetAll().Return([]model.Users{}, nil).Times(1)

	s := service.NewService(mockData)
	users, err := s.Users()

	assert.Equal(t, []model.Users{}, users)
	assert.Nil(t, err)
}

func TestShouldDelegateUser(t *testing.T) {
	mockData := mock.NewMockIUsersData(gomock.NewController(t))

	mockData.EXPECT().GetByUsername("cenk").Return(model.Users{
		Username: "cenk",
		Balance:  155,
	}, nil).Times(1)

	s := service.NewService(mockData)
	users, err := s.User("cenk")

	assert.Equal(t, model.Users{
		Username: "cenk",
		Balance:  155,
	}, users)
	assert.Nil(t, err)
}

func TestShouldDelegateNewUser(t *testing.T) {
	mockData := mock.NewMockIUsersData(gomock.NewController(t))

	user := model.Users{
		Username: "cenk",
		Balance:  155,
	}

	mockData.EXPECT().Insert(model.Users{Username: "cenk", Balance: 155}).Return(nil).Times(1)

	s := service.NewService(mockData)
	err := s.NewUser(user)

	assert.Nil(t, err)
}

func TestShouldDelegateUpdateUser(t *testing.T) {
	mockData := mock.NewMockIUsersData(gomock.NewController(t))

	user := model.Users{
		Username: "cenk",
		Balance:  155,
	}

	mockData.EXPECT().Update(model.Users{Username: "cenk", Balance: 155}).Return(nil).Times(1)

	s := service.NewService(mockData)
	err := s.UpdateUser(user)

	assert.Nil(t, err)
}
