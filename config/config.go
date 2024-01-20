package config

import "log"

type Config struct {
}

var config *Config

var Con = func() *Config {
	if config == nil {
		log.Fatalln("set up  the  config first")
	}

	return config
}
