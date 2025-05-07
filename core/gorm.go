package core

import (
	"gin_web_frame/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
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

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCodePackage{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},
		system.SysParams{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
		example.ExaAttachmentCategory{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	//err = bizModel()

	if err != nil {
		global.LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
