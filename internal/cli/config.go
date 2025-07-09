package cli

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	FilePath  string `mapstructure:"file_path"`
	ServerURL string `mapstructure:"server_url"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка парсинга файла конфигурации: %w", err)
	}

	return &cfg, nil
}
