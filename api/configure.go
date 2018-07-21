package api

import (
	"fmt"
	"log"

	"github.com/burntsushi/toml"
)

var CONFIG *Config

type DBConfig struct {
	Host     string `toml:"host"`
	Dbname   string `toml:"dbname"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	SSLMode  string `toml:"sslmode"`
}

type Config struct {
	Db DBConfig `toml:"database"`
}

// Dhstr returns a Postgresql formatted database string
func (c *Config) Dbstr() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=%s",
		c.Db.Host, c.Db.Dbname, c.Db.User, c.Db.Password, c.Db.Port, c.Db.SSLMode)
}

// initConfig reads a TOML config file into global CONFIG object
func initConfig() {
	var err error
	c := Config{}

	if _, err = toml.DecodeFile(*flagConfigPath, &c); err != nil {
		log.Fatal("error decoding TOML config: ", err)
	}

	CONFIG = &c
}
