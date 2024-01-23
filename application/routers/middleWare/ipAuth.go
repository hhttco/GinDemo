package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hhttco/GinDemo/application/config"
	"github.com/hhttco/GinDemo/application/helpers"
)

// ip 白名单校验
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range config.Config.AllowIp {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			helpers.ResponseError(c, 500, fmt.Sprintf("%v, not in iplist", c.ClientIP()))
			c.Abort()
			return
		}
		c.Next()
	}
}
