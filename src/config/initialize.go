package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const config_file_name = "../config.yaml"
const banner_file_name = "../banner.txt"

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

func PrintBanner() {
	data, err := ioutil.ReadFile(banner_file_name)
	if err != nil {
		log.Fatal("Couldn't load application banner...", err)
	}
	fmt.Println(string(data))
	fmt.Println()
}

type Configuration struct {
	No_No_Words []string `yaml:"great_firewall"`
}
