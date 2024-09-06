package config

import "github.com/spf13/viper"

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func Load(path string) (c Config, err error) {
	viper.AddConfigPath("/home/dunice/Документы/todoGolang")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetConfigFile(path)

	err = viper.ReadInConfig()

	if err != nil {
		return c, err
	}

	err = viper.Unmarshal(&c)

	return c, err
}
