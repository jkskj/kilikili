package api

import (
	"github.com/gin-gonic/gin"
	"kilikili/logics/service"
	"kilikili/util/e"
	"kilikili/util/middleware"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	//绑定结构体
	code := e.SUCCESS
	err := c.ShouldBind(&userRegister)
	if err == nil {
		res := userRegister.Register()
		c.JSON(code, res)
	} else {
		code = e.InvalidParams
		c.JSON(code, err)
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	code := e.SUCCESS
	err := c.ShouldBind(&userLogin)
	if err == nil {
		res := userLogin.Login()
		c.JSON(code, res)
	} else {
		code = e.InvalidParams
		c.JSON(code, err)
	}
}

// UserChangePersonal 用户更改个人信息
func UserChangePersonal(c *gin.Context) {
	var userChange service.UserService
	code := e.SUCCESS
	err := c.ShouldBind(&userChange)
	chaim, _ := middleware.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := userChange.ChangePersonal(chaim.Id)
		c.JSON(code, res)
	} else {
		code = e.InvalidParams
		c.JSON(code, err)
	}
}

// UserChangePassword 用户更改密码
func UserChangePassword(c *gin.Context) {
	var userChange service.UserService
	code := e.SUCCESS
	err := c.ShouldBind(&userChange)
	chaim, _ := middleware.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := userChange.ChangePersonal(chaim.Id)
		c.JSON(code, res)
	} else {
		code = e.InvalidParams
		c.JSON(code, err)
	}
}
