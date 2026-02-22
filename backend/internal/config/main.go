package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var configManager *viper.Viper

func InitConfig() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("could not read config: %w", err)
	}

	return nil
}

func readConfig() error {
	configManager = viper.New()

	configName := strings.Join([]string{"config", os.Getenv("ENVIRONMENT")}, "-")

	configManager.SetConfigName(configName)
	configManager.SetConfigType("yaml")
	configManager.AddConfigPath("./configs")
	configManager.AddConfigPath(".")

	return configManager.ReadInConfig()
}
