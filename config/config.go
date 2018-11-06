package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port            string
	ReadTimeoutSec  int
	WriteTimeoutSec int
	JwtSecret	string
	DemoAccountDump string
}

func Load() *Configuration {

	viper.AddConfigPath("config")
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Can't read config file: %s \n", err))
	}

	var cfg = new(Configuration)
	err = viper.Unmarshal(cfg)

	if err != nil {
		panic(fmt.Errorf("Can't unmarshal config file: %s \n", err))
	}

	return cfg
}
