package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigFile("conf/setting.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("open setting err")
	}
	// fmt.Println(viper.GetString("db.dns"))
}
