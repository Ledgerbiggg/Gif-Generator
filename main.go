package main

import (
	"gif_generator/src/config"
	"gif_generator/src/logs"
	"gif_generator/src/services"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		logs.Module,
		services.Module,
	)
	app.Run()
}
