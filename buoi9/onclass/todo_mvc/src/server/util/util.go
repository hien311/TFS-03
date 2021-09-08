package util

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DBconfig struct {
	DBDriver string `mapstructure:"DBDriver"`
	DSN      string `mapstructure:"DSN"`
}

func LoadDBConfig(path string) (config DBconfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("driver")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}

	err = viper.Unmarshal(&config)
	return
}
