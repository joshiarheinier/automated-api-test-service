package helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-service/testapp/config/apperror"
)

func WriteResponse(httpCode int, res interface{}, c *gin.Context) {
	body, err := json.Marshal(res)
	if err != nil {
		err := apperror.NewV1Error("ERR000", err)
		WriteResponse(err.HttpCode, err, c)
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(httpCode)
	c.Writer.Write(body)
}
