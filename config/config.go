package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Servers []Server `toml:"Servers"`
}

type Server struct {
	Name        string `toml:"name"`
	Address     string `toml:"address"`
	Description string `toml:"description"`
}

func LoadConfig() (Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "go-ssh", "config.toml")

	// Checks if the file exists, if not it will create one
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfig(configPath); err != nil {
			return Config{}, err
		}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading the config file: %v", err)
	}

	var cfg Config
	if _, err := toml.Decode(string(data), &cfg); err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %v", err)
	}

	return cfg, nil
}

func createDefaultConfig(path string) error {
	defaultConfig := Config{
		Servers: []Server{
			{Name: "Raspberry Pi’s", Address: "pi@192.168.1.10"},
			{Name: "Linux", Address: "user@linux-server.com"},
		},
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating config file: %v", err)
	}
	defer file.Close()
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(defaultConfig)
	if err != nil {
		return fmt.Errorf("erro writing default config file: %v", err)

	}
	log.Printf("Default configuration file created at %s. Plese edit it to add your servers.\n", path)
	return nil
}
