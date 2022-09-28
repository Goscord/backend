package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var config *Config = nil

type Config struct {
	Port         int    `json:"port"`
	GithubSecret string `json:"github_secret"`
	Database     struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
}

func Load() error {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		config := new(Config)

		config.Port = 80
		config.Database.Host = "127.0.0.1"

		bytes, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("config.json", bytes, 0644)
		if err != nil {
			return err
		}
	}

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}

	config = new(Config)
	_ = json.Unmarshal(data, &config)

	return nil
}

func Get() *Config {
	return config
}
