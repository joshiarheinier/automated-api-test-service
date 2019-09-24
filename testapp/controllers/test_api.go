package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-service/testapp/docs"
	"github.com/joshia/automated-api-test-service/testapp/testlib"
	"io/ioutil"
)


func StartAPITest(c *gin.Context) {
	ex := testlib.Initiate()
	body, _ := ioutil.ReadAll(c.Request.Body)
	testConfig := docs.New(body)
	url := testConfig.Hostname + testConfig.Port
	for _, schema := range testConfig.Schema {
		storedData := make(map[string]interface{})
		for _, job := range schema.Jobs {
			tmp, _ := job.Body.MarshalJSON()
			stringBody := new(bytes.Buffer)
			json.Compact(stringBody, tmp)
			byteBody := stringBody.Bytes()
			if len(storedData) != 0 {
				ex.CheckHeader(job.Header, storedData)
				ex.CheckHeader(job.Params, storedData)
				byteBody = ex.CheckBody(byteBody, storedData)
			}
			job.Url = ex.SetParams(job.Url, job.Params)
			var validator testlib.TestValidator
			for i := 0; i < job.Try; i++ {
				req, _ := ex.MakeRequest(job.Method, url + job.Url, byteBody)
				ex.SetHeader(req, job.Header)
				validator = ex.Execute(req)
				ex.Flush()
			}
			if !validator.ExpectResponseStatus(job.Expected.Status) || !validator.ExpectBody(job.Expected.Body) {
				msg := "Test failed on:\nTest Title: " + schema.TestTitle + "\nTest Name: " + job.TestName
				reply(500, msg, c)
				return
			}
			if len(job.SaveKeys) != 0 {
				for i := range job.SaveKeys {
					storedData[job.SaveKeys[i]] = validator.ExpectNotNilAndSave(job.SaveKeys[i])
				}
			}
		}
	}
	reply(200, "PASS", c)
}

func reply(status int, message string, c *gin.Context) {
	c.JSON(status, &gin.H{
		"message" : message,
	})
}
