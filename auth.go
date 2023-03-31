package main

import (
	"net/http"
	"os"

	"github.com/5elenay/revoltgo"
	"github.com/gookit/config/v2"
	"github.com/pterm/pterm"
)

func checkLogin() {
	if getLoginToken() != "" {
		return
	}

	Clear()
	pterm.DefaultHeader.WithFullWidth().Println(Config().brandName + " - Log In")
	pterm.Println()
	email, _ := pterm.DefaultInteractiveTextInput.Show(" Email")
	pass, _ := pterm.DefaultInteractiveTextInput.Show(" Password")
	Clear()

	loader, _ := pterm.DefaultSpinner.Start("Logging in...")

	client := revoltgo.Client{
		SelfBot: &revoltgo.SelfBot{
			Email:    email,
			Password: pass,
		},
	}

	client.HTTP = &http.Client{}
	err := client.Auth()
	if err != nil {
		loader.Fail("Failed to login: " + err.Error())
		res, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show("Try again?")
		if res {
			checkLogin()
		} else {
			os.Exit(0)
		}
	}

	config.Set("token", client.SelfBot.SessionToken)
	loader.Success()
}
