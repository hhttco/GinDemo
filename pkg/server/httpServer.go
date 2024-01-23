package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/hhttco/GinDemo/application/config"
	"github.com/hhttco/GinDemo/application/controllers/admin"
	"github.com/hhttco/GinDemo/application/routers"
	"net/http"
	"time"
)

var (
	HttpServerHandler *http.Server
)

func HttpServerRun() {
	// 设置 gin 模式

	gin.SetMode(config.Config.DebugMode)
	r := routers.Init()
	admin.DoInit()

	// 读取服务配置
	HttpServerHandler = &http.Server{
		Addr:           config.Config.Port,
		Handler:        r,
		ReadTimeout:    time.Duration(config.Config.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Config.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(config.Config.MaxHeaderBytes), // 1*2^20 = 1MB
	}

	// 启动协程
	go func() {
		// 服务端请求端口
		log.Printf(" [INFO] HttpServerRun: 127.0.0.1%s\n", config.Config.Port)
		if err := HttpServerHandler.ListenAndServe(); err != nil {
			// 致命错误
			log.Fatalf(" [ERROR] HttpServerRun: 127.0.0.1%s error:%v\n", config.Config.Port, err)

		}
	}()
}

// 服务停止
func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpServerHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
