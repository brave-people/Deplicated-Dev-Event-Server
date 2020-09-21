package init

import "github.com/spf13/viper"

func ViperInit() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
