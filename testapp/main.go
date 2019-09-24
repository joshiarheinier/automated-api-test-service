package main

import (
	// "database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-suite/testapp/controllers"
)


func main()  {
	r := gin.Default()

	controllers.SetUpRouter(r)

	// Listen and Server in 0.0.0.0:8183
	r.Run(":8183")
}