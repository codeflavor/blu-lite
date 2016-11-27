package start

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/golang/glog"

	c "github.com/codeflavor/servops/pkg/config"
)

// createDefaults creates default settings for the application.
func createDefaults(wd string) (*c.AppConfig, error) {
	var cwd string
	var err error

	if wd == "" {
		cwd, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	confDir := filepath.Join(cwd, c.DefaultConfDir)
	confFile := filepath.Join(cwd, c.DefaultConfDir, c.DefaultFile)
	glog.V(0).Infof("Using default username: (%s) and password: (%s). THIS IS NOT RECOMMENDED FOR PRODUCTION!!!", c.DefaultDBUser, c.DefaultDBPassword)

	dbConf := c.NewDBConfig(c.DefaultDBUser, c.DefaultDBPassword, c.DefaultDBURL, c.DefaultDBName)
	conf := c.NewAppConfig(cwd, confFile, confDir, c.DefaultLogLevel, dbConf)

	_, err = os.Stat(conf.ConfDir)
	if os.IsNotExist(err) {
		// maybe extract default permissions to a const.
		err = os.MkdirAll(conf.ConfDir, 0700)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	if err := saveDefaults(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func genDefaultConf(workDir string) (*c.AppConfig, error) {
	conf, err := createDefaults(workDir)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func loadUserConfig(workdir string) (*c.AppConfig, error) {
	conf := &c.AppConfig{}
	return conf, nil
}

func saveDefaults(c *c.AppConfig) error {
	var handler *os.File
	_, err := os.Stat(c.ConfFile)

	if os.IsNotExist(err) {
		handler, err = os.Create(c.ConfFile)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	defer handler.Close()
	encodeConf, err := json.MarshalIndent(c, "", "")
	if err != nil {
		return err
	}

	_, err = handler.Write(encodeConf)
	if err != nil {
		return err
	}
	return nil
}
