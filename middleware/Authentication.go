package middleware

import (
	"net/http"

	"github.com/frozturk/gologin/auth"
	"github.com/gin-gonic/gin"
)

func AuthenticationMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.IsTokenValid(c.Request)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
