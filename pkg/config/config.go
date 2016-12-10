package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type dbDriver string

const (
	// Default application settings.
	DefaultPort       = "8080"
	DefaultHost       = "127.0.0.1"
	DefaultDBHost     = "127.0.0.1"
	DefaultConfFile   = "servops"
	DefaultConfDir    = "config.servops"
	DefaultDBUser     = "user"
	DefaultDBPassword = "password"
	DefaultDBName     = "servops"
	// FIXME: fill this in with a proper default URL.
	DefaultDBURL    = ""
	DefaultLogLevel = 0
)

// LoadConfig loads the configuration specified or generates defaults if a
// current one doesn't exist
func LoadConfig(dir, loglevel string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	var confDir string

	if dir == "" {
		confDir = filepath.Join(cwd, DefaultConfDir)
	} else {
		confDir = filepath.Join(dir, DefaultConfDir)
	}

	viper.AddConfigPath(confDir)
	viper.SetConfigName(DefaultConfFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}
}

// AppConfig holds the application configuration and also encompasses the
// database configuration.
type AppConfig struct {
	WorkDir  string   `json:"workdir"`
	ConfFile string   `json:"configFile"`
	ConfDir  string   `json:"confDir"`
	DBConf   DBConfig `json:"dbConfig"`
	LogLevel int      `json:"logLevel"`
	Port     string   `json:"port"`
	Host     string   `json:"host"`
}

// DBConfig holds information that is necessary to connect to a mongodb
// instance.
type DBConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"URL"`
	DBName   string `json:"dbName"`
}

// NewAppConfig is a helper function for instantiating a new application
// configuration.
func NewAppConfig(workDir, confFile, confDir string, loglevel int, dbConf *DBConfig) *AppConfig {
	return &AppConfig{
		WorkDir:  workDir,
		ConfDir:  confDir,
		ConfFile: confFile,
		DBConf:   *dbConf,
		LogLevel: loglevel,
	}
}

// NewDBConfig holds the information that is necessary to connect to mongodb.
func NewDBConfig(user, pass, mongoDBURL, dbName string) *DBConfig {
	return &DBConfig{
		UserName: user,
		Password: pass,
		DBName:   dbName,
		URL:      mongoDBURL,
	}
}
