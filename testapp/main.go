package main

import (
	// "database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-service/testapp/controllers"
	"github.com/joshia/automated-api-test-service/testapp/middleware/wrapper"
	"net/http"
	"time"
)


func main()  {
	r := gin.Default()

	controllers.SetUpRouter(r)
	// Listen and Server in 0.0.0.0:8183
	s := &http.Server{
		Addr:         ":8183",
		Handler:      wrapper.NewHTTPHandler(r),
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
	}
	s.ListenAndServe()
}