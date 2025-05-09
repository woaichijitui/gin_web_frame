package core

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		global.ACTIVE_DBNAME = &global.CONFIG.Mysql.Dbname
		return GormMysql()
	case "pgsql":
		global.ACTIVE_DBNAME = &global.CONFIG.Pgsql.Dbname
		return GormPgSql()
	case "oracle":
		global.ACTIVE_DBNAME = &global.CONFIG.Oracle.Dbname
		return GormOracle()
	case "mssql":
		global.ACTIVE_DBNAME = &global.CONFIG.Mssql.Dbname
		return GormMssql()
	case "sqlite":
		global.ACTIVE_DBNAME = &global.CONFIG.Sqlite.Dbname
		return GormSqlite()
	default:
		global.ACTIVE_DBNAME = &global.CONFIG.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		models.UserModel{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
