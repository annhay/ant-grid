package inits

import (
	"ant-grid/internal/common/global"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ViperInit() {
	viper.SetConfigFile("../../configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
	err = viper.Unmarshal(&global.AppConf)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
	log.Println("viper 动态配置完成:", global.AppConf)
}
