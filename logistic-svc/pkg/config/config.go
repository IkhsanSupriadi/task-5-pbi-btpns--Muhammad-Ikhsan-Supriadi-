package config

import "github.com/spf13/viper"

type Config struct {
	Port      string `mapstructure:"PORT"`
	DBUrl     string `mapstructure:"DB_URL"`
	DBTestUrl string `mapstructure:"DB_TEST_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	var p string
	if path == "" {
		p = "./pkg/config/envs"
	} else {
		p = path
	}

	viper.AddConfigPath(p)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
