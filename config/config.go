// config/config.go
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	BotToken    string `json:"botToken"`
	BotHost     string `json:"botHost"`
	WebhookPath string `json:"webhookPath"`
}

func ReadConfig(filename string) (Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
