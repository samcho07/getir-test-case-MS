package data

import (
	"log"

	"github.com/BurntSushi/toml"
)

type conf struct {
	MongoDB_Server   string
	MongoDB_Database string
}

func (c *conf) Read() {

	// Making error.
	if _, err := toml.DecodeFile("conf.toml", &c); err != nil {
		log.Fatal(err)
	}
}
