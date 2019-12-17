package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-service/testapp/lib/uuid"
)

func InjectRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		rId := c.Request.FormValue("requestId")
		if rId == "" {
			rId = uuid.NewRequestId()
		}
		c.Request.Header.Add("requestId", rId)
		c.Next()
	}
}
