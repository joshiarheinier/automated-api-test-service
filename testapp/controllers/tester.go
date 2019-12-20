package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-service/testapp/config/apperror"
	"github.com/joshia/automated-api-test-service/testapp/controllers/helpers"
	"github.com/joshia/automated-api-test-service/testapp/dto"
	"github.com/joshia/automated-api-test-service/testapp/models"
)


func StartAPITest(c *gin.Context) {
	tc, err := models.GetTesterComponent(c)
	if err != nil {
		err := apperror.NewV1Error("ERR000", err)
		helpers.WriteResponse(err.HttpCode, err, c)
		return
	}
	if err = models.TestAPI(tc); err != nil {
		err := apperror.NewV1Error("ERR000", err)
		helpers.WriteResponse(err.HttpCode, err, c)
		return
	}
	res := &dto.TestPassResponse{Message:"PASS"}
	helpers.WriteResponse(200, res, c)
}


