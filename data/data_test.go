package data_test

import (
	"bootcamp/data"
	"bootcamp/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataGetByUsernameNotExist(t *testing.T) {

	data := data.NewInMemoryData()

	user := model.Users{
		Username: "cenk", Balance: 10,
	}

	err := data.Insert(user)

	assert.Nil(t, err)

	response, err := data.GetByUsername("ahmet")

	assert.NotNil(t, err)
	assert.Equal(t, "User not found", err.Error())
	assert.Equal(t, model.Users{}, response)

}

func TestDataGetAllNotExist(t *testing.T) {

	var users []model.Users

	data := data.NewInMemoryData()

	response, err := data.GetAll()

	assert.NotNil(t, err)
	assert.Equal(t, "No users found", err.Error())
	assert.Equal(t, users, response)

}

func TestDataGetAll(t *testing.T) {
	data := data.NewInMemoryData()

	users := []model.Users{
		{Username: "cenk", Balance: 0},
	}

	err := data.Insert(users[0])
	assert.Nil(t, err)

	response, err := data.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, users, response)
}

func TestDataInsert(t *testing.T) {

	data := data.NewInMemoryData()

	users := model.Users{
		Username: "cenk", Balance: 0,
	}

	err := data.Insert(users)

	assert.Nil(t, err)

}

func TestDataInsertNotExist(t *testing.T) {

	data := data.NewInMemoryData()

	var user model.Users

	err := data.Insert(user)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not insert user", err.Error())

}

func TestDataUpdateNotExist(t *testing.T) {

	data := data.NewInMemoryData()

	var user model.Users

	err := data.Update(user)

	assert.NotNil(t, err)
	assert.Equal(t, "User not found", err.Error())

}

func TestDataUpdate(t *testing.T) {

	data := data.NewInMemoryData()

	users := model.Users{
		Username: "cenk", Balance: 0,
	}

	err := data.Update(users)

	assert.Nil(t, err)

}

func TestDataGetByUsername(t *testing.T) {

	data := data.NewInMemoryData()

	user := model.Users{
		Username: "cenk", Balance: 0,
	}

	err := data.Insert(user)

	assert.Nil(t, err)

	response, err := data.GetByUsername(user.Username)

	assert.Nil(t, err)
	assert.Equal(t, user, response)

}
