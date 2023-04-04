package service

import (
	"kilikili/model"
	"kilikili/util/e"
	"kilikili/util/serializer"
)

type BulletCommentService struct {
	Content string `form:"content" json:"content"`   //弹幕内容
	AddTime string `form:"add_time" json:"add_time"` //弹幕出现在视频的时间
}

// Create 上传弹幕
func (service *BulletCommentService) Create(uid uint, vid string) serializer.Response {
	var user model.User
	model.DB.First(&user, uid)
	var video model.Video
	model.DB.First(&video, vid)
	code := e.SUCCESS
	bulletComment := model.BulletComment{
		User:    user,
		Uid:     user.ID,
		Video:   video,
		Vid:     video.ID,
		Content: service.Content,
		AddTime: service.AddTime,
	}
	err := model.DB.Create(&bulletComment).Error
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

// List 获取弹幕
func (service *BulletCommentService) List(vid string) serializer.Response {
	var bulletComment []model.BulletComment
	count := 0
	code := e.SUCCESS
	err := model.DB.Model(&model.BulletComment{}).Where("vid=?", vid).Count(&count).Find(&bulletComment).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildBulletComments(bulletComment), uint(count))
}
