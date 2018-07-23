package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/burntsushi/toml"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type Settings struct {
	Host     string `toml:"host"`
	Dbname   string `toml:"dbname"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	SSLMode  string `toml:"sslmode"`
}

type DBConfig struct {
	Prod Settings `toml:"prod"`
	Dev  Settings `toml:"dev"`
}

// Dhstr returns a Postgresql formatted database string
func (s Settings) Dbstr() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=%s",
		s.Host, s.Dbname, s.User, s.Password, s.Port, s.SSLMode)
}

func init() {
	var err error
	c := DBConfig{}

	log.Println("reading db config file...")
	if _, err = toml.DecodeFile(os.Getenv("DBCONFIGPATH"), &c); err != nil {
		log.Fatal("error decoding TOML config: ", err)
	}

	log.Println("connecting to DB...")
	var dbstr string
	cmd := exec.Command("uname")
	if stdout, _ := cmd.Output(); string(stdout[:len(stdout)-1]) == "Darwin" { // running on local mac
		log.Println("using development database")
		dbstr = c.Dev.Dbstr()
	} else { // inside docker bridge network container
		log.Println("using production database")
		dbstr = c.Prod.Dbstr()
	}

	DB, err = sql.Open("postgres", dbstr)
	if err != nil {
		log.Fatal("error opening database connection: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to DB")
}
