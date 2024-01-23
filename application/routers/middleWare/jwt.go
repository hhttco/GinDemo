package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hhttco/GinDemo/application/helpers"
	"github.com/hhttco/GinDemo/pkg/jwt"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			helpers.ResponseError(c, 10008, fmt.Sprintf("Invalid token: %v", token))
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			helpers.ResponseError(c, 10009, fmt.Sprintf("Invalid token: %v", token))
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			helpers.ResponseError(c, 10010, fmt.Sprintf("Invalid token: %v", token))
			c.Abort()
			return
		}

		c.Next()
	}
}
