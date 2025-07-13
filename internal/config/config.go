package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	WordCount  int    `mapstructure:"word_count"`
	Difficulty string `mapstructure:"difficulty"`
	TimeLimit  int    `mapstructure:"time_limit"`
	ShowWPM    bool   `mapstructure:"show_wpm"`
	ShowErrors bool   `mapstructure:"show_errors"`
}

func Load() *Config {
	cfg := &Config{
		WordCount:  20,
		Difficulty: "medium",
		TimeLimit:  60,
		ShowWPM:    true,
		ShowErrors: true,
	}

	viper.SetDefault("word_count", 20)
	viper.SetDefault("difficulty", "medium")
	viper.SetDefault("time_limit", 60)
	viper.SetDefault("show_wpm", true)
	viper.SetDefault("show_errors", true)

	if err := viper.Unmarshal(cfg); err != nil {
		return cfg
	}

	return cfg
}

func (c *Config) Save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(home, ".typeracer.yaml")
	return viper.WriteConfigAs(configPath)
}
