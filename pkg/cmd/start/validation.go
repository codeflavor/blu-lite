package start

import (
	"log"

	"github.com/golang/glog"

	"github.com/codeflavor/servops/pkg/config"
	"github.com/codeflavor/servops/pkg/service"
)

// ValidateUserConfig tries to validate the user specified YAML file and if
// none was specified, then create defaults and run with that
func ValidateUserConfig(confFile, workDir, logLevel string) {
	var appConfig *config.AppConfig
	var err error

	// TODO: refactor this whenever possible
	if confFile != "" {
		glog.V(0).Infof("Trying to load user specified configuration from %s...", confFile)
		appConfig, err = loadUserConfig(confFile)
		if err != nil {
			log.Fatalf("error: An error occured while loading the configuration file %s: (%v)", confFile, err)
		}
	} else {
		glog.V(0).Info("No user specified configuration, generating defaults...")
		appConfig, err = genDefaultConf(workDir)
		if err != nil {
			log.Fatalf("error: An error occured while generating a default config: %v ", err)
		}
		glog.V(0).Info("Configuration loaded, trying to start service...")
	}

	if err := service.StartServer(appConfig); err != nil {
		log.Fatalf("error: An error occured while launching the application services: %v", err)
	}
}
