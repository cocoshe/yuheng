package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	*viper.Viper
}

//默认全局变量
var GlobalConfig *Config

//init
func init() {
	GlobalConfig = &Config{
		viper.New(),
	}
	GlobalConfig.SetConfigName("application")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("./config")

	err := GlobalConfig.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	GlobalConfig.WatchConfig() //自动更新配置
	GlobalConfig.OnConfigChange(func(e fsnotify.Event) {
		err := GlobalConfig.ReadInConfig()
		if err == nil {
			log.Print("config updated")
		}
	})

}
