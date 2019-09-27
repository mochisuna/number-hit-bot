package config

import (
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

// New Config
func New(config *Config, configPath string) error {
	_, err := toml.DecodeFile(configPath, config)
	return err

}
