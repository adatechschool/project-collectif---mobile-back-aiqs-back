package main

import (
	"wave/app"
	"wave/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	//app.Run(":8080")
}
