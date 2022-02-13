package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bootcamp/config"
	"bootcamp/handler"
	"bootcamp/mock"
	"bootcamp/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandlerUserHandler(t *testing.T) {
	t.Run("Get all users", func(t *testing.T) {
		service := mock.NewMockIUsersService(gomock.NewController(t))

		service.EXPECT().Users().Return([]model.Users{
			{
				Username: "Cenk",
				Balance:  100,
			},
			{
				Username: "Ayla",
				Balance:  200,
			},
		}, nil).Times(1)

		handler := handler.NewHandler(service)

		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)

		w := httptest.NewRecorder()

		handler.UserHandler(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json", w.Result().Header.Get("content-type"))

		expectedResBody := []model.Users{}
		err := json.Unmarshal(w.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal err")

		assert.Equal(t, expectedResBody[0].Balance, 100)
	})

	t.Run("Get user with username", func(t *testing.T) {
		service := mock.NewMockIUsersService(gomock.NewController(t))

		service.EXPECT().User("cenk").Return(model.Users{
			Username: "cenk",
			Balance:  155,
		}, nil).Times(1)

		handler := handler.NewHandler(service)

		req := httptest.NewRequest(http.MethodGet, "/cenk", http.NoBody)

		w := httptest.NewRecorder()

		handler.UserHandler(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json", w.Result().Header.Get("content-type"))

		expectedResBody := model.Users{}
		err := json.Unmarshal(w.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")

		assert.Equal(t, expectedResBody.Balance, 155)
	})

	t.Run("should return status not found when users service failed", func(t *testing.T) {
		service := mock.NewMockIUsersService(gomock.NewController(t))

		service.
			EXPECT().
			Users().
			Return([]model.Users{}, errors.New("error occured")).
			Times(1)

		handler := handler.NewHandler(service)
		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		w := httptest.NewRecorder()
		handler.UserHandler(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("should return status not found when user service failed", func(t *testing.T) {
		service := mock.NewMockIUsersService(gomock.NewController(t))

		service.
			EXPECT().
			User("ayla").
			Return(model.Users{}, errors.New("User not found")).
			Times(1)

		handler := handler.NewHandler(service)
		req := httptest.NewRequest(http.MethodGet, "/ayla", http.NoBody)
		w := httptest.NewRecorder()
		handler.UserHandler(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("Create a wallet by username", func(t *testing.T) {
		service := mock.NewMockIUsersService(gomock.NewController(t))

		request := model.Users{
			Username: "aylacenk",
			Balance:  config.Get().InitialBalanceAmount,
		}

		service.EXPECT().User("aylacenk").Return(model.Users{}, errors.New("User not found")).Times(1)

		buf, _ := json.Marshal(request)

		service.EXPECT().NewUser(request).Return(nil).Times(1)

		handler := handler.NewHandler(service)
		req := httptest.NewRequest(http.MethodPut, "/aylacenk", bytes.NewBuffer(buf))
		w := httptest.NewRecorder()
		handler.UserHandler(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("Should return error when user not found", func(t *testing.T) {
		service := mock.NewMockIUsersService(gomock.NewController(t))

		request := model.Users{
			Username: "aylacenk",
			Balance:  22,
		}

		service.EXPECT().User("aylacenk").Return(model.Users{}, errors.New("User not found")).Times(1)

		buf, _ := json.Marshal(request)

		userHandler := handler.NewHandler(service)

		req := httptest.NewRequest(http.MethodPost, "/aylacenk", bytes.NewReader(buf))
		w := httptest.NewRecorder()
		userHandler.UserHandler(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Delete method is not allowed", func(t *testing.T) {
		handler := handler.NewHandler(nil)

		req := httptest.NewRequest(http.MethodDelete, "/", nil)

		w := httptest.NewRecorder()

		handler.UserHandler(w, req)

		assert.Equal(t, http.StatusNotImplemented, w.Result().StatusCode)
	})
}
