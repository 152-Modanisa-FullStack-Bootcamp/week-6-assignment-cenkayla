package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bootcamp/config"
	"bootcamp/model"
	"bootcamp/service"
)

type IHandler interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IUsersService
}

func NewHandler(service service.IUsersService) IHandler {
	return &Handler{service: service}
}

func (c *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	config := config.Get()
	userName := r.URL.Path[1:]

	switch r.Method {
	case http.MethodGet:
		if userName == "" {
			response, err := c.service.Users()
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
				return
			}

			json, _ := json.Marshal(&response)

			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(json)
			return
		}
		response, err := c.service.User(userName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		json, _ := json.Marshal(&response)

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)

	case http.MethodPost:
		user, err := c.service.User(userName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("User not exists"))
			return
		}

		var req struct {
			Balance int `json:"balance"`
		}

		json.NewDecoder(r.Body).Decode(&req)

		user.Balance += req.Balance
		if user.Balance < config.MinumumBalanceAmount {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Balance cannot be lower than " + fmt.Sprintf("%d", config.MinumumBalanceAmount)))
			return
		}
		err = c.service.UpdateUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Add("content-type", "application/json")

		w.WriteHeader(http.StatusOK)

	case http.MethodPut:
		check, _ := c.service.User(userName)
		if check.Username != "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("User already exists"))
			return
		}
		user := model.Users{
			Username: userName,
			Balance:  config.InitialBalanceAmount,
		}

		err := c.service.NewUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusNotImplemented)
	}
}
