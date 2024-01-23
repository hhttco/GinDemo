package middleWare

import (
	"github.com/gin-gonic/gin"
)

var (
	corsAllowHeaders     = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// Access-Control-Allow-Origin 跨域请求
func AccessControlAllowOrigin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Credentials", corsAllowCredentials)
		c.Header("Access-Control-Allow-Headers", corsAllowHeaders)
		c.Header("Access-Control-Allow-Methods", corsAllowMethods)
		c.Header("Access-Control-Allow-Origin", corsAllowOrigin)

		c.Next()
	}
}
