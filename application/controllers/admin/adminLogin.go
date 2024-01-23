package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hhttco/GinDemo/application/controllers"
	"github.com/hhttco/GinDemo/application/helpers"
	"github.com/hhttco/GinDemo/application/models/admin"
	"github.com/hhttco/GinDemo/application/routers"
	"github.com/hhttco/GinDemo/pkg/jwt"
)

func init() {
	vc := LoginViewCtrl{}
	// 测试mvc页面
	// r.GET("/admin/index", vc.Index)
	routers.NoTokenApiGet("/admin/index", vc.Index)
	// 登录 生产 JWT
	routers.NoTokenApiPost("/admin/login", vc.Login)
}

type LoginViewCtrl struct {
}

/**
 * 获取用户信息
 *
 */
func (vc *LoginViewCtrl) Index(c *gin.Context) {
	//panic(1)
	//helpers.ResponseOk(c, "测试")
	new(admin.Admin).TableName()
	controllers.LoadAdminView(c, "admin/index.html", gin.H{
		"tplName": "admin/index.html",
		"title":   "登录",
	})
}

/**
 * 管理员登录
 *
 */
func (vc *LoginViewCtrl) Login(c *gin.Context) {
	//panic(1)
	adminId := jwt.GetAdminId(c)
	if len(adminId) <= 0 {
		// 未登录
		token, err := jwt.GenerateToken("112", "张三", "mi123")
		if err != nil {
			return
		}

		helpers.ResponseOk(c, token)
		return
	}

	helpers.ResponseOk(c, "用户已登录")
}
