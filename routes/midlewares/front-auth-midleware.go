package midlewares

import (
	"github.com/gin-gonic/gin"
	"solar-faza/utils"
	"strings"
)

func UserAuthMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate logged in user
		cookie := c.GetHeader("Cookie")
		for _, cookieSlice := range strings.Split(cookie, ";") {
			if strings.Contains(cookieSlice, "data") {
				token := strings.Split(cookieSlice, "=")[1]
				c.Set("userId", utils.GetDataFromToken(token))
			}
		}
		c.Next()
	}
}
