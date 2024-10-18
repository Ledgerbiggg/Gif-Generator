package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var config *GConfig

// GConfig 配置
type GConfig struct {
	Mod      string `yaml:"mod"`      // 模式，如 "dev"
	LogLevel int    `yaml:"logLevel"` // 日志级别

	Output struct {
		GifName string `yaml:"gifName"` // 生成的 GIF 文件名
	} `yaml:"output"`

	Image struct {
		Width  int `yaml:"width"`  // GIF 宽度
		Height int `yaml:"height"` // GIF 高度

		// 修改 imageList 为包含文件名和延迟的结构体数组
		ImageList []struct {
			File  string `yaml:"file"`  // 图像文件名
			Delay int    `yaml:"delay"` // 每帧的延迟（毫秒）
		} `yaml:"imageList"`
	} `yaml:"image"`

	GifSettings struct {
		Loop    int `yaml:"loop"`    // 循环次数（0 = 无限循环）
		Quality int `yaml:"quality"` // GIF 质量（数值越低质量越好，体积越大）
	} `yaml:"gifSettings"`

	Options struct {
		Dither       bool `yaml:"dither"`       // 是否启用抖动
		ResizeImages bool `yaml:"resizeImages"` // 是否调整图像大小
	} `yaml:"options"`
}

func LoadConfig() *GConfig {
	vconfig := viper.New()

	vconfig.AutomaticEnv()
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	vconfig.SetConfigName("config")
	vconfig.AddConfigPath(".")
	vconfig.SetConfigType("yaml")

	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := vconfig.Unmarshal(&config); err != nil {
		log.Panicln("\"unmarshal cng file fail " + err.Error())
	}
	return config
}
