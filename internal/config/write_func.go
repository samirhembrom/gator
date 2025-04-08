package config

import (
	"encoding/json"
	"os"
	"os/user"
)

func SetUser(config Config) error {
	err := write(&config)
	if err != nil {
		return err
	}
	byte, err := json.Marshal(config)
	if err != nil {
		return err
	}
	address, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(address, byte, 0644)
	if err != nil {
		return err
	}

	return nil
}

func write(cfg *Config) error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}
	username := currentUser.Username
	cfg.CurrentUser = username
	return nil
}
