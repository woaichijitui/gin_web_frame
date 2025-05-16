package core

import (
	"fmt"
	"gin_web_frame/config"
	"gin_web_frame/core/internal"
	"gin_web_frame/global"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ByteZhou-2018](https://github.com/ByteZhou-2018)
func GormMysql() *gorm.DB {
	m := global.CONFIG.Mysql
	return initMysqlDatabase(m)
}

// GormMysqlByConfig 通过传入配置初始化Mysql数据库
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	return initMysqlDatabase(m)
}

// initMysqlDatabase 初始化Mysql数据库的辅助函数
func initMysqlDatabase(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.LOG.Info(fmt.Sprintf("mysql connect success: %s", m.Dsn()))
		return db

	}
}
