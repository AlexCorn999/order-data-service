package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DataBaseURL string `toml:"database_url"`
	ClusterID   string `toml:"cluster_id"`
	ClientID    string `toml:"client_id"`
}

const configPath = "./configs/apiserver.toml"

func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}

// ParseFlags parses flags from the toml file into the Config structure at the specified path.
func (c *Config) ParseFlags() {
	flag.Parse()

	if _, err := toml.DecodeFile(configPath, c); err != nil {
		log.Fatal(err)
	}
}
