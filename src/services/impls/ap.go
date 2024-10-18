package impls

import (
	"gif_generator/src/config"
	"gif_generator/src/logs"
	"go.uber.org/fx"
)

// Params params for all
type Params struct {
	fx.In
	LC     fx.Lifecycle
	Log    *logs.ConsoleLogger
	Config *config.GConfig
}
