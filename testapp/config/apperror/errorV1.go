package apperror

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"strconv"
)

const (
	ErrFailedToDecodeConfigurationFile = "Failed to decode configuration file: %v\n"
)

type V1Error struct {
	ErrorCode		string	`json:"error_code"`
	PublicMessage	string	`json:"public_message"`
	DebugMessage	string	`json:"debug_message"`
	HttpCode		int		`json:"http_code"`
}

var errorData  = make(map[string]*V1Error)

func (e *V1Error) Error() string {
	errorStr := "public_message:" + e.PublicMessage + ";" +
		"debug_message:" + e.DebugMessage + ";" +
		"http_code:" + strconv.Itoa(e.HttpCode)
	return errorStr
}

func init()  {
	_, filename, _, _ := runtime.Caller(0)
	errorConfigFilePath := path.Join(path.Dir(filename), "./json/errorV1.json")
	errorConfigFile, _ := ioutil.ReadFile(errorConfigFilePath)
	err := json.Unmarshal(errorConfigFile, &errorData)
	if err != nil {
		log.Fatalf(ErrFailedToDecodeConfigurationFile, err)
	}
}

func NewV1Error(code string, args ...interface{}) *V1Error {
	err := errorData[code]
	err.DebugMessage = fmt.Sprintf(err.DebugMessage, args...)
	return err
}