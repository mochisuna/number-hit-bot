package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Config all settings
type Config struct {
	Server    `toml:"server"`
	LineBot   `toml:"linebot"`
	Firestore `toml:"firestore"`
}

// Server port
type Server struct {
	Port string `toml:"port"`
}

type Firestore struct {
	ProjectID string `toml:"project_id"`
	JSONPath  string `toml:"json_path"`
}

// LineBot
type LineBot struct {
	ChannelToken  string `toml:"channel_token"`
	ChannelSecret string `toml:"channel_secret"`
}

// New create config
func New(env string) (*Config, error) {
	configPath := fmt.Sprintf("_config/%s/config.toml", env)
	config := &Config{}
	_, err := toml.DecodeFile(configPath, config)
	return config, err
}
