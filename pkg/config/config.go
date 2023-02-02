package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(c interface{}) error {
	// main viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// read the file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// map to app
	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	// done
	log.Printf("config loaded successfully")
	return nil
}
