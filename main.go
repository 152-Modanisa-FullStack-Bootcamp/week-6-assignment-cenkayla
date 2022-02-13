package main

import (
	"fmt"
	"net/http"

	"bootcamp/config"
	"bootcamp/data"
	"bootcamp/handler"
	"bootcamp/service"
)

func main() {

	config := config.Get()

	data := data.NewInMemoryData()
	service := service.NewService(data)
	handler := handler.NewHandler(service)

	http.HandleFunc("/", handler.UserHandler)

	err := http.ListenAndServe(config.ServerAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
