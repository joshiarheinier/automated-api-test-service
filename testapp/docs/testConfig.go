package docs

import (
	"encoding/json"
	"log"
)

const (
	ErrFailedToDecodeConfigurationFile = "Failed to decode configuration file: %v\n"
)

type TestConfiguration struct {
	ConfigName	string			`json:"config_name"`
	Hostname	string			`json:"hostname"`
	Port		string			`json:"port"`
	Schema		[]TestSchema	`json:"schema"`
}

type TestSchema struct {
	TestTitle		string				`json:"test_title"`
	Jobs			[]SubTestConfig		`json:"jobs"`
}

type SubTestConfig struct {
	TestName		string				`json:"test_name"`
	Try				int					`json:"try"`
	Url				string				`json:"url"`
	Method			string				`json:"method"`
	Params			map[string]string	`json:"params"`
	Header			map[string]string	`json:"header"`
	Body			json.RawMessage		`json:"body"`
	SaveKeys		[]string			`json:"save_keys"`
	Expected		ExpectedResult		`json:"expected"`
}

type ExpectedResult struct {
	Status	string					`json:"response_status"`
	Body	map[string]interface{}	`json:"body"`
}

var TestConfigData  = &TestConfiguration{}


func New(body []byte) *TestConfiguration {
	err := json.Unmarshal(body, &TestConfigData)
	if err != nil {
		log.Fatalf(ErrFailedToDecodeConfigurationFile, err)
	}
	return TestConfigData
}