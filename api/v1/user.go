package v1

import (
	"MiMengCore/model"
	"MiMengCore/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var reg model.UserRegistration
	// 尝试将请求体绑定到reg变量
	if err := c.ShouldBindJSON(&reg); err != nil {
		// 如果绑定失败，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	userID := reg.UserID
	userName := reg.UserName
	password := reg.Password
	qq := reg.QQ

	// 校验用户ID
	valid, msg := service.CheckUserID(userID)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}

	// 校验用户名
	if !service.CheckUserName(userName) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名格式不正确",
		})
		return
	}

	// 校验密码
	if !service.CheckPassword(password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码格式不正确",
		})
		return
	}

	// 校验QQ
	if !service.CheckQQ(qq) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "QQ格式不正确",
		})
		return
	}

	user := model.User{
		UserID:   userID,
		UserName: userName,
		Password: password,
		QQ:       qq,
		VipDate:  time.Now(),
		IsAdmin:  false,
		Points:   0,
	}

	err := service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "注册失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(c *gin.Context) {
	var login model.LoginCredentials
	// 尝试将请求体绑定到login变量
	if err := c.ShouldBindJSON(&login); err != nil {
		// 如果绑定失败，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	userID := login.UserID
	password := login.Password

	// 校验用户名和密码
	user, valid, msg := service.CheckUserLogin(userID, password)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}

	// 如果用户存在，生成登录令牌
	token, err := service.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "登录令牌生成失败",
		})
		return
	}

	// 返回成功响应和令牌
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": token,
	})
}
