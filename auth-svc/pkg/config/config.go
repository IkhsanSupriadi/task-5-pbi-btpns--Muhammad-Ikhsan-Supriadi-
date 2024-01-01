package config

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	TestPort     string `mapstructure:"TEST_PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	DBTestUrl    string `mapstructure:"DB_TEST_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig(configPath string) (config Config, err error) {
	var cp string
	if configPath == "" {
		cp = "./pkg/config/envs"
	} else {
		cp = configPath
	}
	viper.AddConfigPath(cp)
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
