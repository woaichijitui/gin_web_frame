package global

import (
	"fmt"
	"gin_web_frame/config"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	DB        *gorm.DB
	DBList    map[string]*gorm.DB
	REDIS     redis.UniversalClient
	REDISList map[string]redis.UniversalClient
	MONGO     *qmgo.QmgoClient
	CONFIG    config.Server
	VP        *viper.Viper
	// LOG    *oplogging.Logger
	LOG *zap.Logger
	//Timer               timer.Timer = timer.NewTimerTask()
	//Concurrency_Control             = &singleflight.Group{}
	//ROUTERS             gin.RoutesInfo
	ACTIVE_DBNAME *string
	//BlackCache          local_cache.Cache
	lock sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetRedis(name string) redis.UniversalClient {
	redis, ok := REDISList[name]
	if !ok || redis == nil {
		panic(fmt.Sprintf("redis `%s` no init", name))
	}
	return redis
}
