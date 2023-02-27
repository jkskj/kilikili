package api

import (
	"github.com/gin-gonic/gin"
	"kilikili/logics/service"
	"kilikili/util/e"
	"kilikili/util/middleware"
)

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
