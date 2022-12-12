package main

type Configuration struct {
	brandName string
	dev       bool
}

func Config() Configuration {
	return Configuration{
		brandName: "CLIvolt",
		dev:       true,
	}
}
