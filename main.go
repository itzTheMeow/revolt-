package main

import (
	"github.com/5elenay/revoltgo"
	"github.com/pterm/pterm"
)

func main() {
	Clear()

	pterm.Println("Loading settings...")
	initSettings()

	pterm.Println("Checking login...")
	checkLogin()
	Clear()

	// A BasicText printer is used to print text, without special formatting.
	// As it implements the TextPrinter interface, you can use it in combination with other printers.
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")
	pterm.DefaultBasicText.Println("For example to resolve progressbars and spinners.")
	// If you just want to print text, you should use this instead:
	// 	pterm.Println("Hello, World!")

	pterm.DefaultBasicText.Println("test")
	pterm.DefaultBasicText.Println(getLoginToken())

	client := revoltgo.Client{
		SelfBot: &revoltgo.SelfBot{
			SessionToken: getLoginToken(),
			//UserId:       "-",
		},
	}

	client.OnMessage(func(m *revoltgo.Message) {
		pterm.DefaultBasicText.Println(m.Content)
	})
}
