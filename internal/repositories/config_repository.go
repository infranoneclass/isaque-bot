package repositories

import (
	"encoding/json"
)

type Config struct {
	BotToken          string
	HoneypotOnlyKicks bool `json:"honeypot_only_kicks"`
}

func NewConfigRepository() *Config {
	return &Config{}
}

func (c *Config) BindFromJSONFile(file []byte) error {
	return json.Unmarshal(file, &c)
}

func (c *Config) BindBotToken(token string) {
	c.BotToken = token
}

func (c *Config) GetIfHoneypotOnlyKicks() bool {
	return c.HoneypotOnlyKicks
}
