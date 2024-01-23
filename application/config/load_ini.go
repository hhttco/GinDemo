package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

var Config *config

type config struct {
	Address        string
	Port           string
	AddressPort    string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
	AllowIp        []string
	source         *ini.File

	// base
	DebugMode string

	// log
	LogLevel string

	LogFileWriterOn              bool
	LogFileWriterLogPath         string
	LogFileWriterRotateLogPath   string
	LogFileWriterWfLogPath       string
	LogFileWriterRotateWfLogPath string

	LogConsoleWriterOn    bool
	LogConsoleWriterColor bool

	// jwt
	JwtSecret string
}

/**
 * 判断配置文件是否存在
 *
 */
//func (c *config) Load(path string) *config {
//	var err error
//	exists, err := helpers.PathExists(path)
//	if !exists {
//		fmt.Println("配置文件不存在")
//		return c
//	}
//
//	c.source, err = ini.Load(path)
//	if err != nil {
//		fmt.Println("配置文件加载失败")
//		panic(err)
//	}
//
//	return c
//}

/**
 * 判断配置是否加载成功
 *
 */
func (c *config) Init(source *ini.File) *config {
	c.source = source

	if c.source == nil {
		return c
	}

	c.Address = c.source.Section("server").Key("ADDRESS").MustString("127.0.0.1")
	c.Port = fmt.Sprintf(":%d", c.source.Section("server").Key("PORT").MustInt(8013))
	c.ReadTimeout = c.source.Section("server").Key("READ_TIMEOUT").MustInt(10)
	c.WriteTimeout = c.source.Section("server").Key("WRITE_TIMEOUT").MustInt(10)
	c.MaxHeaderBytes = c.source.Section("server").Key("MAX_HEADER_BYTES").MustInt(20)
	c.AllowIp = c.source.Section("server").Key("ALLOW_IP").Strings(",")

	c.DebugMode = c.source.Section("base").Key("DEBUG_MODE").MustString("debug")

	c.LogLevel = c.source.Section("log").Key("LOG_LEVEL").MustString("trace")
	c.LogFileWriterOn = c.source.Section("log.file_writer").Key("ON").MustBool(true)
	c.LogFileWriterLogPath = c.source.Section("log.file_writer").Key("LOG_PATH").MustString("./logs/gin_scaffold.inf.log")
	c.LogFileWriterRotateLogPath = c.source.Section("log.file_writer").Key("ROTATE_LOG_PATH").MustString("./logs/gin_scaffold.inf.log.%Y%M%D%H")
	c.LogFileWriterWfLogPath = c.source.Section("log.file_writer").Key("WF_LOG_PATH").MustString("./logs/gin_scaffold.wf.log")
	c.LogFileWriterRotateWfLogPath = c.source.Section("log.file_writer").Key("ROTATE_WF_LOG_PATH").MustString("./logs/gin_scaffold.wf.log.%Y%M%D%H")

	c.LogConsoleWriterOn = c.source.Section("log.console_writer").Key("ON").MustBool(false)
	c.LogConsoleWriterColor = c.source.Section("log.console_writer").Key("COLOR").MustBool(false)

	c.JwtSecret = c.source.Section("jwt").Key("JWT_SECRET").MustString("")

	return c
}
