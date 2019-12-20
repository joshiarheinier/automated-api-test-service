package config

import (
	"encoding/json"
	"github.com/joshia/automated-api-test-service/testapp/config/apperror"
	"github.com/joshia/automated-api-test-service/testapp/config/structs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

type Config struct {
	Logger	*config.LogConfiguration `json:"logger"`
}

var configData  = &Config{}

func init()  {
	var ok bool
	env, ok := os.LookupEnv("GOENV")
	if !ok {
		env = "development"
	}
	_, filename, _, _ := runtime.Caller(0)
	configFilePath := path.Join(path.Dir(filename), "./json/"+env+"/settings.json")
	configFile, _ := ioutil.ReadFile(configFilePath)
	err := json.Unmarshal(configFile, &configData)
	if err != nil {
		log.Fatalf(apperror.ErrFailedToDecodeConfigurationFile, err)
	}
}

func NewConfig() *Config {
	return configData
}