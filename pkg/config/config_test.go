package config

import (
	"testing"

	"reflect"

	"gopkg.in/mgo.v2"
)

const (
	workDir  = "/test"
	confFile = "/test/servops.yaml"
	confDir  = "/test/config.servops"
	loglevel = 5

	user = "test1"
	pass = "test1"
	url  = "mongo/test"
)

var (
	session = &mgo.Session{}
)

func newTestDBConfig() *DBConfig {
	return &DBConfig{
		UserName: user,
		Password: pass,
		URL:      url,
	}
}

func TestNewAppConfig(t *testing.T) {
	dbConf := newTestDBConfig()
	testAppConf := &AppConfig{
		WorkDir:  workDir,
		ConfDir:  confDir,
		ConfFile: confFile,
		DBConf:   *dbConf,
		LogLevel: loglevel,
	}
	AppConfInst := NewAppConfig(workDir, confFile, confDir, loglevel, dbConf)
	if !reflect.DeepEqual(testAppConf, AppConfInst) {
		t.Errorf("Expected new instance to match: %#v, got %#v", testAppConf, AppConfInst)
	}
}
