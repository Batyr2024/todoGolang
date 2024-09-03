package config

import "github.com/spf13/viper"

type Config struct{
	Port string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`

}

func LoadConfig() (c Config, err error){
	viper.AddConfigPath("./pfg/common/config")
}