package middleware

import (
	"net/http"

	"github.com/frozturk/jwtauth"
	"github.com/gin-gonic/gin"
)

func AuthenticationMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtauth.IsAuthenticated(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.Next()
	}
}
