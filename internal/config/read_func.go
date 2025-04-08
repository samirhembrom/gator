package config

import (
	"encoding/json"
	"io"
	"os"
)

func Read() (Config, error) {
	address, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	jsonFile, err := os.Open(address)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/" + configFileName, nil
}
