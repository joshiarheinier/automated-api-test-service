package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {

	// API test route
	r.POST("/v1/api/test", StartAPITest)
}
