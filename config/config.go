package config

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Config struct {
	ServerAddr           string `json:"serverAddr"`
	InitialBalanceAmount int    `json:"initialBalanceAmount"`
	MinumumBalanceAmount int    `json:"minumumBalanceAmount"`
}

var c = &Config{}

func init() {
	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if strings.Contains(mydir, "handler") {
		os.Chdir("..")
	}

	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, c)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return c
}
