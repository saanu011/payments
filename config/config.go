package config

import (
	"github.com/spf13/viper"
	"log"
	"payments/pkg/database"
	"payments/pkg/server"
)

type Config struct {
	Database database.Config
	Server   server.Config
}

func Load(configFile string) (*Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	v.SetConfigName(configFile)
	v.SetConfigType("yml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to read config file: %v\n", err)
		return nil, err
	}

	var config Config
	unmarshalErr := v.Unmarshal(&config)

	return &config, unmarshalErr
}
