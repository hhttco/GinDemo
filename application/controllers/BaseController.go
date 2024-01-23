package controllers

import "github.com/gin-gonic/gin"

/**
 * 加载模版
 */
func LoadAdminView(c *gin.Context, tplName string, data interface{}) {
	c.HTML(200, tplName, data)
}

/**
 * 加载模版
 */
func LoadView(c *gin.Context, data interface{}) {
	c.HTML(200, "public/user_dash", data)
}
