package inits

import "ant-grid/internal/common/global"

// GatewayInit 网关层初始化
func GatewayInit() {
	UserGrpcInit()
}

// ServerInit 服务初始化
func ServerInit() {
	ViperInit() //viper 初始化
	// 初始化Zap日志
	global.Logger = InitLogger()
	defer global.Logger.Sync() // 延迟关闭日志，确保日志写入
	MysqlInit()                //mysql 初始化
	RedisInit()                //redis 初始化
}
