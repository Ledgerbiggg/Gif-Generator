package services

import (
	"gif_generator/src/logs"
	"gif_generator/src/services/impls"
	"gif_generator/src/services/interfaces"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(impls.NewGenGif),
	fx.Invoke(func(i interfaces.IGenGif, logger *logs.ConsoleLogger) {
		logger.Info("service init")
		err := i.GenerateGif()
		if err != nil {
			logger.Error(err.Error())
		}
	}),
)
