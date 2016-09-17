package start

import (
	"github.com/codeflavor/servops/pkg/config"
)

const (
	// we only use the defaults if the user didn't specify any?
	confFileName = "servops.yml"
	confDirName  = "servops"
)

// createDefaults creates default settings for the application.
func createDefaults(workDir string) (*config.AppConfig, error) {
	conf := &config.AppConfig{}
	return conf, nil
}

func generateDefaultConfig(workDir string) (*config.AppConfig, error) {
	conf := &config.AppConfig{}
	return conf, nil
}

func loadUserConfig(workdir string) (*config.AppConfig, error) {
	conf := &config.AppConfig{}
	return conf, nil
}
