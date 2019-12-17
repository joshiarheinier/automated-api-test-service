package config

import (
	"encoding/json"
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

var testConfigData  = &TestConfiguration{}


func NewTestConfig(body []byte) (*TestConfiguration, error) {
	err := json.Unmarshal(body, &testConfigData)
	if err != nil {
		return nil, err
	}
	return testConfigData, nil
}