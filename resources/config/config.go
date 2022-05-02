package config

import (
	"fmt"
	"os"
)

// Config is the struct to define our applciation configuration
type Config struct {
	Port     string         `json:"port"`
	Env      string         `json:"env"`
	Database PostgresConfig `json:"database"`
}

// PostgresConfig is the struct to define our postgres configuration
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

// Dialect states that we are utilizing postgres
func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// Connection makes the db connection string
func (c PostgresConfig) Connection() string {
	if c.Host == "localhost" {
		if c.Password == "" {
			return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Dbname)
		}
		return fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", c.User, c.Password, c.Host, c.Dbname)
	}
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s", c.User, c.Password, c.Host, c.Dbname)
}

// LoadConfig initializes the db based on env variables
func LoadConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" || os.Getenv("ENVIRONMENT") == "dev" {
		setTheSecrets()
	}

	var (
		c    Config
		host = os.Getenv("DBHOST")
		name = os.Getenv("DBNAME")
		user = os.Getenv("DBUSER")
		pass = os.Getenv("DBPASS")
		port = os.Getenv("PORT")
		env  = os.Getenv("ENVIRONMENT")
	)

	c.Database.Host = host
	c.Database.Dbname = name
	c.Database.User = user
	c.Database.Password = pass
	c.Database.Port = 5432

	c.Port = port
	c.Env = env

	return c
}
