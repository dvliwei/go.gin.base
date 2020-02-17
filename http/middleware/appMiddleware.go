package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func AppMiddleWare() gin.HandlerFunc{

	return func(context *gin.Context) {
		t :=time.Now()
		context.Set("example", "12345")
		context.Next()
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := context.Writer.Status()
		log.Println(status)
	}
}