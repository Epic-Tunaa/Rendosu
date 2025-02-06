package config

import (
	"os"
	"path/filepath"
	"github.com/spf13/viper"
)

type Config struct {
	Audio struct {
		SampleRate	int
		BufferSize	int
	}
	Video struct {
		Fullscreen	bool
		Resolution	[2]int
		VSync	bool
	}
	Input struct {
		KeyBindings map[string]glfw.Key
	}
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		setupDefaultConfig()
	}
}

func setupDefaultConfig() {
	viper.SetDefault("audio.samplerate", 44100)
	viper.SetDefault("video.resolution", [2]int{1920, 1080})
	viper.SetDefault("input.keybindings", map[string]int{
		"pause": int(glfw.KeySpace)
		"skip": int(glfw.KeyRight)
	})

	if err := viper.SafeWriteConfig(); err != nil {
		log.Fatal("Failed to create default config", err)
	}
}

func Get() *Config {
	var cfg Config
	viper.Unmarshall(&cfg)
	return &cfg
}