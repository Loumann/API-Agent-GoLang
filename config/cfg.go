package config

import (
	"encoding/json"
	"log"
	"os"
)

var configPatch = "config/cfg.json"

type Config struct {
	Host     string `json:"host" binding:"required"`
	Port     string `json:"port" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Dbname   string `json:"dbname" binding:"required"`
	SSLmode  string `json:"ssl_mode" binding:"required"`
}

func ReadCfg() (*Config, error) {
	var config Config

	file, err := os.ReadFile(configPatch)
	if err != nil {
		log.Fatalln("error reading config file")
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln("error parsing config file")
	}
	return &config, err
}
