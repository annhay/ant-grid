package router

import (
	"ant-grid/internal/api-gateway/handler"
	"ant-grid/internal/api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	public := router.Group("/api/v1/public")
	{
		public.POST("/sendTextMessage", handler.SendTextMessage) //短信发送
		public.POST("/register", handler.Register)               //注册
		public.POST("/login", handler.Login)                     //登录
	}

	return router
}
