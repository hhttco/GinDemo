package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/hhttco/GinDemo/application/config"
	"github.com/hhttco/GinDemo/pkg/logging"
	"sync"
	"time"
)

/**
 * 初始化默认连接引擎
 * config.Mysql 设置成数组即可实现多库 主从
 *
 */

var (
	GORMMapPool        map[string]*gorm.DB
	_defaultEngineName string
	__once             sync.Once
)

func init() {
	__once.Do(func() {
		GORMMapPool = make(map[string]*gorm.DB, 0)
	})

	if err := InitDBPool(); err != nil {
		logging.Error("mysql初始化失败", err)
		return
	}
}

func InitDBPool() error {
	// 如果 config 是数组 则循环实现不同库的连接池
	db, err := gorm.Open("mysql", config.Mysql.DataSourceName)
	if err != nil {
		return err
	}

	db.SingularTable(true)
	err = db.DB().Ping()
	if err != nil {
		return err
	}

	db.LogMode(true)
	//db.SetLogger()

	db.DB().SetMaxIdleConns(config.Mysql.MaxIdleConn)
	db.DB().SetMaxOpenConns(config.Mysql.MaxOpenConn)
	db.DB().SetConnMaxLifetime(time.Duration(config.Mysql.MaxConnLifeTime) * time.Second)

	GORMMapPool[config.Mysql.DriverName] = db
	_defaultEngineName = config.Mysql.DriverName

	return nil
}

// 通过 DRIVER_NAME 查找
func GetPoolEngine(name string) (*gorm.DB, error) {
	if engine, ok := GORMMapPool[name]; ok {
		return engine, nil
	}

	return nil, errors.New("get pool error")
}

// 获取
func GetDbEngine() *gorm.DB {
	if engine, ok := GORMMapPool[_defaultEngineName]; ok {
		return engine
	}

	return nil
}

func CloseDB() error {
	for _, dbPool := range GORMMapPool {
		dbPool.Close()
	}
	return nil
}

// 表结构同步
func DoSynBeans(bean interface{}) {
	engine := GetDbEngine()
	if !engine.HasTable(bean) {
		engine.AutoMigrate(bean)
	}
}
