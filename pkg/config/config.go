package config

type dbDriver string

const (
	// Default application settings.
	DefaultPort       = "8080"
	DefaultHost       = "127.0.0.1"
	DefaultDBHost     = "127.0.0.1"
	DefaultConfDir    = "config.servops"
	DefaultFile       = "servops.yaml"
	DefaultDBUser     = "user"
	DefaultDBPassword = "password"
	DefaultDBName     = "servops"
	// FIXME: fill this in with a proper default URL.
	DefaultDBURL    = ""
	DefaultLogLevel = 0
)

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
