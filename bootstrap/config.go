package bootstrap

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	g "Raven-gin/global"
)

func InitializeConfig() *viper.Viper {
	config := "config.yaml"
	if configEnv := os.Getenv("RAVEN_CONFIG"); configEnv != "" {
		config = configEnv
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %w", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		if err := v.Unmarshal(&g.Cof); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&g.Cof); err != nil {
		fmt.Println(err)
	}

	return v

}
