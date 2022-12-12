package main

import "github.com/pterm/pterm"

func checkLogin() {
	if getLoginToken() != "" {
		return
	}

	Clear()
	pterm.DefaultHeader.WithFullWidth().Println(Config().brandName + " - Log In")
	pterm.Println()
	res, _ := pterm.DefaultInteractiveTextInput.WithTextStyle(pterm.NewStyle()).Show(" Email")
	pterm.Println(res)
	res2, _ := pterm.DefaultInteractiveTextInput.Show(" Password")
	pterm.Println(res2)
}
