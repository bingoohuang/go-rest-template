package middlewares

import "github.com/gin-gonic/gin"

// CORS middleware
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		f := c.Writer.Header().Set
		f("Access-Control-Allow-Origin", "*")
		f("Access-Control-Allow-Credentials", "true")
		f("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		f("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}
