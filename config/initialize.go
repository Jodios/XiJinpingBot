package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const config_file_name = "config.yaml"

func InitContext(config *Configuration) {
	f, err := os.Open(config_file_name)
	if err != nil {
		log.Fatal("Couldn't load application context...", err)
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		log.Fatal("Couldn't load application context...", err)
	}
}

type Configuration struct {
	Token string `yaml:"bot_token"`
}
