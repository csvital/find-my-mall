package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read(filePath string) {
	if _, err := toml.DecodeFile(filePath, &c); err != nil {
		log.Fatal(err)
	}
}
