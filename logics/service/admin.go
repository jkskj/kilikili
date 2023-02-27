package service

import (
	"kilikili/model"
	"kilikili/util/e"
	"kilikili/util/serializer"
)

type AdminService struct {
	Uid         uint  `form:"uid"json:"uid"`
	Vid         uint  `form:"vid" json:"vid"`
	Pass        int64 `form:"pass" json:"pass"`
	CommentType int64 `form:"comment_type"json:"comment_type"`
	Cid         int64 `form:"cid" json:"cid"`
}

func (service *AdminService) Block() serializer.Response {
	var user model.User
	var count int
	code := e.SUCCESS
	model.DB.Model(&model.User{}).First(&user, service.Uid).Count(&count)
	if count == 0 {
		code = e.ErrorNotExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.IsBlock = 1
	model.DB.Save(user)
	return serializer.Response{
		Status: code,
		Msg:    "已拉黑用户",
	}
}
func (service *AdminService) Examine() serializer.Response {
	var video model.Video
	var user model.User
	code := e.SUCCESS
	err := model.DB.First(&video, service.Vid).Error
	model.DB.First(&user, video.Uid)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	video.User = user
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildVideo(video),
		Msg:    e.GetMsg(code),
	}
}
func (service *AdminService) IsPass() serializer.Response {
	var video model.Video
	var user model.User
	code := e.SUCCESS
	err := model.DB.First(&video, service.Vid).Error
	model.DB.First(&user, video.Uid)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	video.IsPass = service.Pass
	err = model.DB.Save(video).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	video.User = user
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildVideo(video),
		Msg:    e.GetMsg(code),
	}
}
func (service *AdminService) Delete() serializer.Response {
	var err error
	code := e.SUCCESS
	if service.CommentType == 0 {
		var bulletComment model.BulletComment
		err = model.DB.First(&bulletComment, service.Cid).Error
		model.DB.Delete(bulletComment)
	} else if service.CommentType == 1 {
		var comment model.Comment
		err = model.DB.First(&comment, service.Cid).Error
		model.DB.Delete(comment)
	} else {
		var reply model.Reply
		err = model.DB.First(&reply, service.Cid).Error
		model.DB.Delete(reply)
	}
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
