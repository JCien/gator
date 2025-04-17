package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(homeDir, configFileName)
	return path, nil

}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	
	data, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer data.Close()

	decoder := json.NewDecoder(data)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
