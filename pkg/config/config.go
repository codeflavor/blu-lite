package config

type dbDriver string

// AppConfig holds the application configuration and also encompasses the
// database configuration.
type AppConfig struct {
	WorkDir  string   `json:"workdir"`
	ConfFile string   `json:"configfile"`
	DbConf   dbConfig `json:"dbconfig"`
	LogLevel string   `json:"loglevel"`
}

// dbConfig holds information that is necessary to connect to a mongodb
// instance.
type dbConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"URL"`
}

// newAppConfig is a helper function for instantiating a new application
// configuration.
func newAppConfig(workDir, confFile, loglevel string, dbConf *dbConfig) *AppConfig {
	return &AppConfig{
		WorkDir:  workDir,
		ConfFile: confFile,
		DbConf:   *dbConf,
		LogLevel: loglevel,
	}
}

// newDbConfig holds the information that is necessary to connect to mongodb.
func newDbConfig(user, pass, mongodbURL string) *dbConfig {
	return &dbConfig{
		UserName: user,
		Password: pass,
		URL:      mongodbURL,
	}
}

// generateDefaults generates a default configuration for the application when
// no application configuration file was specified by the user.
func (app *AppConfig) generateDefaults() error {
	return nil
}
