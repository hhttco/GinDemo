package config

import (
	"github.com/go-ini/ini"
)

var Mysql *mysql

type mysql struct {
	DriverName      string
	DataSourceName  string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifeTime int
}

/**
 * 判断配置是否加载成功
 *
 */
func (m *mysql) Init(source *ini.File) *mysql {
	if source == nil {
		return m
	}

	m.DriverName = source.Section("mysql").Key("DRIVER_NAME").MustString("mysql")
	m.DataSourceName = source.Section("mysql").Key("DATA_SOURCE_NAME").MustString("root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	m.MaxOpenConn = source.Section("mysql").Key("MAX_OPEN_CONN").MustInt(20)
	m.MaxIdleConn = source.Section("mysql").Key("MAX_IDLE_CONN").MustInt(10)
	m.MaxConnLifeTime = source.Section("mysql").Key("MAX_CONN_LIFE_TIME").MustInt(100)

	return m
}
