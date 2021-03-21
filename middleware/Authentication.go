package middleware

import (
	"net/http"

	"github.com/frozturk/gotodo/auth"
	"github.com/gin-gonic/gin"
)

func AuthenticationMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.IsAuthenticated(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.Next()
	}
}
