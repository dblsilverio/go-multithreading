package configs

import "github.com/spf13/viper"

type config struct {
	Timeout int `mapstructure:"API_TIMEOUT"`
}

func LoadConfig(path string) (*config, error) {
	var cfg *config

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
