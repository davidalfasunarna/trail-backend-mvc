package core

import "os"

// Config struct
type Config struct {
	DBHost     string
	DBName     string
	DBUsername string
	DBPassword string
	DBDriver   string
	DBPort     string
	DBSslMode  string
}

// Fetch from env variables
func (c *Config) Fetch() {
	c.DBHost = os.Getenv("DB_HOST")
	c.DBName = os.Getenv("DB_NAME")
	c.DBUsername = os.Getenv("DB_USERNAME")
	c.DBPassword = os.Getenv("DB_PASSWORD")
	c.DBDriver = os.Getenv("DB_DRIVER")
	c.DBPort = os.Getenv("DB_PORT")
	c.DBSslMode = os.Getenv("DB_SSL_MODE")
}
