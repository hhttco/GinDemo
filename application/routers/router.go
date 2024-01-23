package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hhttco/GinDemo/application/controllers/user"
	"github.com/hhttco/GinDemo/application/routers/middleWare"
)

var (
	// map[分组名][方式 POST GET][路由地址]方法名
	_routerMap = make(map[string]map[string]map[string]gin.HandlerFunc, 0)

	// map[方式 POST GET][路由地址]方法名
	_routerMapNoToken = make(map[string]map[string]gin.HandlerFunc, 0)
)

func init() {
	if _routerMapNoToken["POST"] == nil ||
		_routerMapNoToken["GET"] == nil {
		_routerMapNoToken["POST"] = make(map[string]gin.HandlerFunc, 0)
		_routerMapNoToken["GET"] = make(map[string]gin.HandlerFunc, 0)
	}
}

func Init(middleWares ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middleWares...)
	// r.Use(middleWare.IPAuthMiddleware()) // ip限制
	r.Use(middleWare.RecoveryMiddleWare())
	r.Use(middleWare.AccessControlAllowOrigin())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 加载静态资源
	r.Static("/static", "./static") //加载静态资源

	// 加载模版文件
	r.LoadHTMLGlob("application/views/**/*")

	/******** admin ********/
	// 路由分组
	// 非登录页面和接口

	user.IndexUserInit(r)

	// 路由注册
	RegisterNoTokenRouter(r)

	r.Use(middleWare.JWT())
	RegisterRouter(r)

	return r
}

/**
 * 路由注册
 */
func RegisterRouter(r *gin.Engine) {
	for g, p := range _routerMap {
		// 判断路由组使用不同的中间件
		// r.Use()

		// 注册
		if g == "" {
			for rType, hList := range p {
				for path, h := range hList {
					switch rType {
					case "POST":
						r.POST(path, h)
					case "GET":
						r.GET(path, h)
					default:
						r.POST(path, h)
					}
				}
			}
		}

		if g != "" {
			gName := r.Group(g)
			for rType, hList := range p {
				for path, h := range hList {
					switch rType {
					case "POST":
						gName.POST(path, h)
					case "GET":
						gName.GET(path, h)
					default:
						gName.POST(path, h)
					}
				}
			}
		}
	}
}

/**
 * 路由注册
 */
func RegisterNoTokenRouter(r *gin.Engine) {
	for t, p := range _routerMapNoToken {
		for path, h := range p {
			switch t {
			case "POST":
				r.POST(path, h)
			case "GET":
				r.GET(path, h)
			default:
				r.POST(path, h)
			}
		}
	}
}

/**
 * 初始化路由 map
 */
func initRouterMap(groupName string) {
	if _routerMap[groupName] == nil {
		_routerMap[groupName] = make(map[string]map[string]gin.HandlerFunc, 0)
		_routerMap[groupName]["POST"] = make(map[string]gin.HandlerFunc, 0)
		_routerMap[groupName]["GET"] = make(map[string]gin.HandlerFunc, 0)
	}
}

/**
 * POST 分组名称 路径 方法
 */
func ApiPost(groupName, path string, h gin.HandlerFunc) {
	initRouterMap(groupName)
	_routerMap[groupName]["POST"][path] = h
}

/**
 * GET 分组名称 路径 方法
 */
func ApiGet(groupName, path string, h gin.HandlerFunc) {
	initRouterMap(groupName)
	_routerMap[groupName]["GET"][path] = h
}

/**
 * 不校验 token POST 路径 方法
 */
func NoTokenApiPost(path string, h gin.HandlerFunc) {
	_routerMapNoToken["POST"][path] = h
}

/**
 * 不校验 token GET 分组名称 路径 方法
 */
func NoTokenApiGet(path string, h gin.HandlerFunc) {
	_routerMapNoToken["GET"][path] = h
}
