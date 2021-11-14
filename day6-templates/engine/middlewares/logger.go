package middlewares

import (
	"gee/engine"
	"log"
	"time"
)

func Logger() engine.HandlerFunc {
	return func(c *engine.Context) {
		t := time.Now()
		// 等待用户自己定义的 Handler处理结束后，再做一些额外的操作
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
