package helpers

import (
	"github.com/gin-gonic/gin"
)

func ResponseOk(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "",
		"data": data,
	})
}

func ResponseError(c *gin.Context, code int, msg interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
}
