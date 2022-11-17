package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Port    string `yaml:"port"`
	CsvFile string `yaml:"csvFile"`
	PokeUrl string `yaml:"pokeUrl"`
}

func LoadConfig() (AppConfig, error) {
	config := AppConfig{}

	yamlFile, err := os.ReadFile("configs.yaml")
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
