package main

import (
	"github.com/hhttco/GinDemo/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server.HttpServerRun()

	// 开启优雅关停 使用 channel 进行线程同步问题
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	// 将数据 传给 channel
	<-quit

	// 服务器停止
	server.HttpServerStop()
}
