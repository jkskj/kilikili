package service

import (
	"github.com/jinzhu/gorm"
	"kilikili/model"
	"kilikili/util/e"
	"kilikili/util/middleware"
	"kilikili/util/serializer"
)

type UserService struct {
	UserName    string `form:"user_name" json:"user_name" `
	Password    string `form:"password" json:"password" `
	Email       string `form:"email" json:"email" `
	Avatar      string `form:"avatar" json:"avatar"`
	NewPassword string `form:"new_password" json:"new_password"`
	Role        int64  `form:"role"json:"role"`
}

// Register 注册方法
func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	code := e.SUCCESS
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	//如果用户已存在
	if count == 1 {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	user.Email = service.Email
	user.Avatar = service.Avatar
	user.Role = service.Role
	//加密密码
	err := user.SetPassword(service.Password)
	if err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err1 := model.DB.Create(&user).Error
	if err1 != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err1.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 登陆方法
func (service *UserService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	err := model.DB.Where("user_name=?", service.UserName).First(&user).Error
	if err != nil {
		//用户是否存在
		if gorm.IsRecordNotFoundError(err) {
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				//Error : err.Error(),
			}
		}
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if user.IsBlock == 1 {
		return serializer.Response{
			Status: code,
			Msg:    "账号已被封",
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err1 := middleware.GenerateToken(user.ID, user.UserName, int(user.Role))
	if err1 != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err1.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: e.GetMsg(code),
	}
}
func (service *UserService) ChangePersonal(uid uint) serializer.Response {
	code := e.SUCCESS
	var user model.User
	err := model.DB.Model(&model.User{}).Where("id=?", uid).First(&user).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.UserName = service.UserName
	user.Email = service.Email
	user.Avatar = service.Avatar
	model.DB.Save(&user)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) ChangePassword(uid uint) serializer.Response {
	code := e.SUCCESS
	var user model.User
	err := model.DB.Model(&model.User{}).Where("id=?", uid).First(&user).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err1 := user.SetPassword(service.NewPassword)
	if err1 != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err1.Error(),
		}
	}
	model.DB.Save(&user)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
