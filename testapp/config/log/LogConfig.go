package config

import (
	"encoding/json"
	"github.com/joshia/automated-api-test-service/testapp/errors"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

type LogConfiguration struct {
	FilePath			string	`json:"file_path"`
	LogFileName			string	`json:"log_file_name"`
	ErrorLogFileName	string	`json:"error_log_file_name"`
	DebugLogFileName	string	`json:"debug_log_file_name"`
	MaxSize				int		`json:"max_size"`
	MaxBackup			int		`json:"max_backup"`
	MaxAge				int		`json:"max_age"`
	Compress			bool	`json:"compress"`
	Level				string	`json:"level"`
}

var logConfigData  = &LogConfiguration{}

func init()  {
	_, filename, _, _ := runtime.Caller(0)
	logConfigFilePath := path.Join(path.Dir(filename), "./json/settings.json")
	logConfigFile, _ := ioutil.ReadFile(logConfigFilePath)
	err := json.Unmarshal(logConfigFile, &logConfigData)
	if err != nil {
		log.Fatalf(errors.ErrFailedToDecodeConfigurationFile, err)
	}
}

func NewLogConfig() *LogConfiguration {
	return logConfigData
}