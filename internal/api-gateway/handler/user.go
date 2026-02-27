package handler

import (
	_const "ant-grid/internal/api-gateway/const"
	"ant-grid/internal/api-gateway/middleware"
	"ant-grid/internal/api-gateway/request"
	"ant-grid/internal/common/global"
	"ant-grid/internal/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendTextMessage 短信发送
func SendTextMessage(c *gin.Context) {
	var form request.SendTextMessage
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := global.UserGrpc.SendTextMessage(c, &proto.SendTextMessageReq{Phone: form.Phone})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "短信发送成功"})
}

// Register 注册
func Register(c *gin.Context) {
	var form request.Register
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	register, err := global.UserGrpc.Register(c, &proto.RegisterReq{
		Phone:            form.Phone,
		VerificationCode: form.VerificationCode,
		Password:         form.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token := middleware.CreateToken(uint64(register.Id), _const.JwtKey)
	c.JSON(http.StatusOK, gin.H{"message": "注册成功", "token": token})
}

// Login 登录
func Login(c *gin.Context) {
	var form request.Login
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	login, err := global.UserGrpc.Login(c, &proto.LoginReq{
		Phone:            form.Phone,
		Password:         form.Password,
		VerificationCode: form.VerificationCode,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token := middleware.CreateToken(uint64(login.Id), _const.JwtKey)
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": token})
}
