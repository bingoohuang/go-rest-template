package middlewares

import (
	"net/http"

	"github.com/bingoohuang/go-rest-template/pkg/crypto"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("authorization")
		if crypto.ValidateToken(h) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}
