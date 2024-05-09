package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of our YAML configuration.
type Config struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

func main() {
	// YAML string containing the configuration data.
	yamlData := `
host: example.com
port: 8080
debug: true
`

	// Declare a variable of type Config.
	var config Config

	// Unmarshal the YAML into our struct. Here, if there is an error during unmarshaling,
	// it will be caught and logged.
	err := yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Output the unmarshaled configuration.
	fmt.Printf("Config: %+v\n", config)
}
