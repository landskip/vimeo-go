package vimeo

import (
	"github.com/BurntSushi/toml"
	"errors"
)

var filePath = "config.toml"

// Root Configuration
type Config struct {
	BaseURL     string `toml:"base_url"`
	AccessToken string `toml:"access_token"`
}

// Create new config object
func NewConfig() (Config, error) {
	var config Config
	_, err := toml.DecodeFile(filePath, &config)

	if config.AccessToken == "" || config.AccessToken == "YOUR TOKEN" {
		return config, errors.New(
			"Required access_token is missing from config.toml")
	}
	return config, err
}
