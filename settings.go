package main

import (
	"os"
	"path"
	"strings"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
)

func initSettings() {
	config.AddDriver(json.Driver)

	dir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	if Config().dev {
		dir, _ = os.Getwd()
	}

	dir = path.Join(dir, strings.ToLower(Config().brandName))
	_, err = os.Stat(dir)
	if err != nil {
		os.Mkdir(dir, os.ModePerm)
	}
	dir = path.Join(dir, "config.json")
	_, err = os.Stat(dir)
	if err != nil {
		fil, _ := os.Create(dir)
		fil.WriteString("{}")
		fil.Close()
	}

	config.WithOptions(config.ParseEnv, config.SaveFileOnSet(dir, "json"))
	err = config.LoadFiles(dir)
	if err != nil {
		panic(err)
	}
}

func getLoginToken() string {
	return config.String("token")
}
