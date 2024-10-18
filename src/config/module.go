package config

import (
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(LoadConfig),
	fx.Invoke(func(c *GConfig) {
		fmt.Println("config: ", c)
	}),
)
