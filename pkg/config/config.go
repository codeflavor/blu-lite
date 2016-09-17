package config

import (
	"gopkg.in/mgo.v2"
)

type dbDriver string

// AppConfig holds the application configuration and also encompasses the
// database configuration.
type AppConfig struct {
	WorkDir  string   `json:"workdir"`
	ConfFile string   `json:"configfile"`
	DbConf   dbConfig `json:"dbconfig"`
	LogLevel string   `json:"loglevel"`
}

type dbConfig struct {
	UserName string    `json:"username"`
	Password string    `json:"password"`
	URL      string    `json:"URL"`
	Driver   dbDriver  `json:"driver"`
	Session  dbSession `json:"session"`
}

// DbHandler holds the active connection to the database
type dbSession struct {
	mongoSession *mgo.Session
}

// newAppConfig is a helper function for instantiating a new application
// configuration
func newAppConfig() *AppConfig {
	return &AppConfig{}
}

func newDbConfig(user, pass, mongodbURL string) *dbConfig {
	return &dbConfig{
		UserName: user,
		Password: pass,
		URL:      mongodbURL,
	}
}

func (app *AppConfig) generateDefaults() error {
	return nil
}
