package config

import (
	"testing"

	"reflect"

	"gopkg.in/mgo.v2"
)

const (
	workDir  = "/test"
	confFile = "/test/servops.yaml"
	loglevel = "5"

	user = "test1"
	pass = "test1"
	url  = "mongo/test"
)

var (
	session = &mgo.Session{}
)

func newTestDBConfig() *dbConfig {
	return &dbConfig{
		UserName: user,
		Password: pass,
		URL:      url,
	}
}

func TestNewAppConfig(t *testing.T) {
	dbConfig := newTestDBConfig()
	testAppConfig := &AppConfig{
		WorkDir:  workDir,
		ConfFile: confFile,
		DbConf:   *dbConfig,
		LogLevel: loglevel,
	}
	AppConfInstance := newAppConfig(workDir, confFile, loglevel, dbConfig)
	if !reflect.DeepEqual(testAppConfig, AppConfInstance) {
		t.Errorf("Expected new instance to match: %#v, got %#v", testAppConfig, AppConfInstance)
	}
}
