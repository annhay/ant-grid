package global

import (
	"ant-grid/internal/common/config"
	"ant-grid/internal/proto"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	AppConf  *config.AppConfig
	DB       *gorm.DB
	Rdb      *redis.Client
	Logger   *zap.Logger
	UserGrpc proto.UserClient
)
