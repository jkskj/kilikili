package service

import (
	"kilikili/model"
	"kilikili/util/e"
	"kilikili/util/serializer"
	"time"
)

type VideoService struct {
	Title   string `form:"title" json:"title" `
	Content string `form:"content" json:"content"`
}

func (service *VideoService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	video := model.Video{
		User:    user,
		Uid:     user.ID,
		Title:   service.Title,
		Content: service.Content,
		Uptime:  time.Now().Unix(),
	}
	code := e.SUCCESS
	//数据库储存任务
	err := model.DB.Create(&video).Error
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
		Data:   serializer.BuildVideo(video),
		Msg:    e.GetMsg(code),
	}
}
func (service *VideoService) Watch(vid string) serializer.Response {
	var video model.Video
	var user model.User
	code := e.SUCCESS
	err := model.DB.First(&video, vid).Error
	model.DB.First(&user, video.Uid)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if video.IsPass != 1 {
		return serializer.Response{
			Status: code,
			Msg:    "视频未通过审核",
		}
	}
	video.User = user
	video.AddView() //增加点击数
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildVideo(video),
		Msg:    e.GetMsg(code),
	}
}
