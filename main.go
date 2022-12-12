package main

import (
	"github.com/5elenay/revoltgo"
	"github.com/pterm/pterm"
)

func main() {
	initSettings()

	// A BasicText printer is used to print text, without special formatting.
	// As it implements the TextPrinter interface, you can use it in combination with other printers.
	pterm.DefaultBasicText.Println("Default basic text printer.")
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")
	pterm.DefaultBasicText.Println("For example to resolve progressbars and spinners.")
	// If you just want to print text, you should use this instead:
	// 	pterm.Println("Hello, World!")

	pterm.DefaultBasicText.Println(getLoginToken())

	client := revoltgo.Client{
		SelfBot: &revoltgo.SelfBot{
			Id:           "session id",
			SessionToken: "session token",
			UserId:       "user id",
		},
	}

	client.OnMessage(func(m *revoltgo.Message) {
		pterm.DefaultBasicText.Println(m.Content)
	})
}
