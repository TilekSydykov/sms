package midlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solar-faza/service"
	"strings"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		accessToken := strings.Split(header, " ")
		if len(accessToken) < 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := service.ValidateToken(accessToken[1])
		if err != nil && !user.IsSuperuser {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		accessToken := strings.Split(header, " ")
		if len(accessToken) < 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := service.ValidateToken(accessToken[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
	}
}

func PublicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		accessToken := strings.Split(header, " ")
		if len(accessToken) > 1 {
			user, err := service.ValidateToken(accessToken[1])
			if err == nil {
				c.Set("user", user)
			}
		}
	}
}
