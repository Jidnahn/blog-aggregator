package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
	Connection        string `json:"connection"`
}

func (c *Config) SetUser(userName string) error {
	c.Current_user_name = userName
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error encoding username: %w", err)
	}
	os.WriteFile(path, data, 0200)

	return nil
}

func Read() (*Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return &Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, fmt.Errorf("error reading file: %w", err)
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return &Config{}, fmt.Errorf("error decoding file: %w", err)
	}

	return &config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}
	path := home + configFileName
	return path, nil
}
