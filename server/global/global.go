package global

import (
	"go.uber.org/zap"
	"sync"

	"gin-vue-admin/config"
	"github.com/antlabs/timer"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG     *zap.Logger
	TIMEWHEEL   timer.Timer
	APP_MANAGER  sync.Map
)
