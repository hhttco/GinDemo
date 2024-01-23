package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hhttco/GinDemo/application/controllers"
)

func IndexUserInit(r *gin.Engine) {
	vc := IndexViewCtrl{}
	r.GET("/hello", vc.GetUser)
}

type IndexViewCtrl struct {
}

/**
 * 获取用户信息
 *
 */
func (vc *IndexViewCtrl) GetUser(c *gin.Context) {
	//helpers.ResponseOk(c, "测试")
	controllers.LoadView(c, gin.H{
		"tplName": "template user/index.html",
		"title":   "test",
	})
}
