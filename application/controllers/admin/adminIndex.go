package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hhttco/GinDemo/application/helpers"
	"github.com/hhttco/GinDemo/application/routers"
	"github.com/hhttco/GinDemo/pkg/jwt"
)

func init() {
	vc := IndexViewCtrl{}
	// 首页
	routers.ApiPost("v1", "/admin/index", vc.Index)
}

type IndexViewCtrl struct {
}

/**
 * 首页
 *
 */
func (vc *IndexViewCtrl) Index(c *gin.Context) {
	//panic(1)
	adminId := jwt.GetAdminId(c)

	helpers.ResponseOk(c, gin.H{
		"admin_id": adminId,
		"data":     "测试",
	})
}
