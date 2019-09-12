// arquivo de configuração, ele irá ler o dados do arquivo config.toml

package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
