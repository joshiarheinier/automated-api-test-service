package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-service/testapp/config/apperror"
	"github.com/joshia/automated-api-test-service/testapp/models"
)


func StartAPITest(c *gin.Context) {
	tc, err := models.GetTesterComponent(c)
	if err != nil {
		err := apperror.NewV1Error("ERR000", err)
		reply(err.HttpCode, err, c)
		return
	}
	if err := models.TestAPI(tc); err != nil {
		err := apperror.NewV1Error("ERR000", err)
		reply(err.HttpCode, err, c)
		return
	}
	reply(200, gin.H{"message":"PASS"}, c)
}

func reply(status int, res interface{}, c *gin.Context) {
	c.JSON(status, &res)
}
