package config

import (
	"github.com/golang/glog"
	"log"
	"os"

	"github.com/spf13/viper"
)

type dbDriver string

const (
	// Default application settings.
	DefaultPort       = "8080"
	DefaultHost       = "127.0.0.1"
	DefaultDBHost     = "127.0.0.1"
	DefaultCfgFile    = "servops"
	DefaultCfgDir     = "config.servops"
	DefaultDBUser     = "user"
	DefaultDBPassword = "password"
	DefaultDBName     = "servops"
	// FIXME: fill this in with a proper default URL.
	DefaultDBURL    = ""
	DefaultLogLevel = 0
)

// LoadCfg loads the configuration specified or generates defaults if a current
// one doesn't exist
func LoadCfg(dir, loglevel string) (*AppCfg, error) {
	var confDir string

	if dir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		confDir = cwd
	} else {
		confDir = dir
	}

	viper.AddConfigPath(confDir)
	viper.SetConfigName(DefaultCfgFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c AppCfg
	viper.Unmarshal(&c)

	return &c, nil
}

// AppCfg holds the application configuration and also encompasses the database
// configuration.
type AppCfg struct {
	WorkDir  string `json:"workdir"`
	ConfFile string `json:"configFile"`
	ConfDir  string `json:"confDir"`
	DBConf   DBCfg  `json:"dbConfig"`
	LogLevel int    `json:"logLevel"`
	Port     string `json:"port"`
	Host     string `json:"host"`
}

// DBCfg holds information that is necessary to connect to a mongodb instance.
type DBCfg struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"URL"`
	DBName   string `json:"dbName"`
}

// NewAppCfg is a helper function for instantiating a new application
// configuration.
func NewAppCfg(workDir, confFile, confDir string, loglevel int, dbConf *DBCfg) *AppCfg {
	return &AppCfg{
		WorkDir:  workDir,
		ConfDir:  confDir,
		ConfFile: confFile,
		DBConf:   *dbConf,
		LogLevel: loglevel,
	}
}

// NewDBCfg holds the information that is necessary to connect to mongodb.
func NewDBCfg(user, pass, mongoDBURL, dbName string) *DBCfg {
	return &DBCfg{
		UserName: user,
		Password: pass,
		DBName:   dbName,
		URL:      mongoDBURL,
	}
}

func setDefaultCfg() {
	viper.SetDefault("port", DefaultPort)
	viper.SetDefault("host", DefaultHost)
	viper.SetDefault("loglevel", DefaultLogLevel)
}
