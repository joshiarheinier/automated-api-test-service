package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
	"testing"
)


func TestAPIFlow(t *testing.T)  {
	client := http.Client{}
	_, filename, _, _ := runtime.Caller(0)
	testConfigFilePath := path.Join(path.Dir(filename), "./docs/json/test/testConfig.json")
	testConfigFile, _ := ioutil.ReadFile(testConfigFilePath)
	req, _ := http.NewRequest("POST", "http://localhost:8183/v1/api/test", bytes.NewReader(testConfigFile))
	res, _ := client.Do(req)
	assert.Equal(t, 200, res.StatusCode)
	byteBody, _ := ioutil.ReadAll(res.Body)
	body := make(map[string]string)
	json.Unmarshal(byteBody, &body)
	assert.Equal(t, "PASS", body["message"])
}
