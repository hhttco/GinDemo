package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hhttco/GinDemo/application/config"
	"github.com/hhttco/GinDemo/application/helpers"
	"github.com/hhttco/GinDemo/pkg/logging"
)

// RecoveryMiddleWare 捕获所有panic，并且返回错误信息
func RecoveryMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 先做一下日志记录
				logging.Error(err)
				if config.Config.DebugMode != "debug" {
					helpers.ResponseError(c, 500, "系统内部错误")
					return
				} else {
					helpers.ResponseError(c, 500, fmt.Sprintf("系统内部错误: %s", err))
					return
				}
			}
		}()
		c.Next()
	}
}
