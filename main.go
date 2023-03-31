package main

import (
	"os"
	"os/signal"
	"syscall"

	revoltgo "github.com/itzTheMeow/revolt-go"
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

	client := revoltgo.Client{
		SelfBot: &revoltgo.SelfBot{
			SessionToken: getLoginToken(),
		},
	}

	client.OnMessage(func(m *revoltgo.Message) {
		pterm.DefaultBasicText.Println(m.Content)
	})

	client.Start()

	// wait for close
	sc := make(chan os.Signal, 1)

	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)
	<-sc

	client.Destroy()
}
