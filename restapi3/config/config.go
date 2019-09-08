package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config criando uma struct para passarmos os dados de conexão do nosso db
type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}