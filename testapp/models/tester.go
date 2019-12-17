package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	config "github.com/joshia/automated-api-test-service/testapp/config/tester"
	"github.com/joshia/automated-api-test-service/testapp/lib/testlib"
	"io/ioutil"
)

type TesterComponent struct {
	Executor	*testlib.TestExecutor
	Config		*config.TestConfiguration
	Hostname	string
}


func GetTesterComponent(c *gin.Context) (*TesterComponent, error) {
	ex := testlib.Initiate()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	testConfig, err := config.NewTestConfig(body)
	if err != nil {
		return nil, err
	}
	hostname := testConfig.Hostname + testConfig.Port
	tc := &TesterComponent{
		Executor: ex,
		Config:   testConfig,
		Hostname: hostname,
	}
	return tc, nil
}

func TestAPI(tc *TesterComponent) error {
	for _, schema := range tc.Config.Schema {
		storedData := make(map[string]interface{})
		for _, job := range schema.Jobs {
			tmp, _ := job.Body.MarshalJSON()
			stringBody := new(bytes.Buffer)
			json.Compact(stringBody, tmp)
			byteBody := stringBody.Bytes()
			if len(storedData) != 0 {
				tc.Executor.CheckHeader(job.Header, storedData)
				tc.Executor.CheckHeader(job.Params, storedData)
				byteBody = tc.Executor.CheckBody(byteBody, storedData)
			}
			job.Url = tc.Executor.SetParams(job.Url, job.Params)
			var validator testlib.TestValidator
			for i := 0; i < job.Try; i++ {
				req, _ := tc.Executor.MakeRequest(job.Method, tc.Hostname + job.Url, byteBody)
				tc.Executor.SetHeader(req, job.Header)
				validator = tc.Executor.Execute(req)
				tc.Executor.Flush()
			}
			if !validator.ExpectResponseStatus(job.Expected.Status) || !validator.ExpectBody(job.Expected.Body) {
				msg := "Test failed on:\nTest Title: " + schema.TestTitle + "\nTest Name: " + job.TestName
				return errors.New(msg)
			}
			if len(job.SaveKeys) != 0 {
				for i := range job.SaveKeys {
					storedData[job.SaveKeys[i]] = validator.ExpectNotNilAndSave(job.SaveKeys[i])
				}
			}
		}
	}
	return nil
}