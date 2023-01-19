package main

import (
	"github/madi-api/app"
	"github/madi-api/config"
)

func main() {
	config := config.NewConfig()

	app := new(app.App)
	app.Initialize(config)
	app.Run(config.ServerHost)
}
