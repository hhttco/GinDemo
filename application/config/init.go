package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/hhttco/GinDemo/application/helpers"
	"github.com/hhttco/GinDemo/pkg/logging"
	"os"
)

func init() {
	// 配置文件路径 判断环境 开发环境 生产环境
	separator := helpers.GetPathSeparator()
	iniPath := helpers.GetProjectPath() + separator + "application" + separator + "config" + separator + "app.ini"

	exists, err := helpers.PathExists(iniPath)
	if err != nil {
		fmt.Println("配置文件不存在", err)
		logging.Error("配置文件不存在", err)
		os.Exit(-1)
	}

	if !exists {
		fmt.Println("配置文件不存在")
		logging.Error("配置文件不存在")
		os.Exit(-1)
	}

	source, err := ini.Load(iniPath)
	if err != nil {
		fmt.Println("配置文件加载失败")
		panic(err)
	}

	// 服务配置
	Config = (&config{}).Init(source)

	// mysql 配置
	Mysql = (&mysql{}).Init(source)
}
