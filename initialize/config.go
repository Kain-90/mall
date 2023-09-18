package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"mall/global"
	"os"
)

func SetupConfig() {
	_, err := os.Stat(global.ConfigFile)
	if err != nil {
		panic(fmt.Sprintf("%s not found", global.ConfigFile))
	}
	v := viper.New()
	v.SetConfigFile(global.ConfigFile)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.String())
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(err)
	}
}
