package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-service/testapp/models"
)


func StartAPITest(c *gin.Context) {
	tc, err := models.GetTesterComponent(c)
	if err != nil {
		reply(500, err.Error(), c)
	}
	if err := models.TestAPI(tc); err != nil {
		reply(500, err.Error(), c)
	}
	reply(200, "PASS", c)
}

func reply(status int, message string, c *gin.Context) {
	c.JSON(status, &gin.H{
		"message" : message,
	})
}
