package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cfgFile, err := os.Open(path.Join(homeDir, configFileName))
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()
	var cfg Config
	decoder := json.NewDecoder(cfgFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) SetUser(user string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cfgFile, err := os.OpenFile(path.Join(homeDir, configFileName), os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer cfgFile.Close()

	cfg.CurrentUserName = user

	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = cfgFile.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
