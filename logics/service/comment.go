package service

import (
	"kilikili/model"
	"kilikili/util/e"
	"kilikili/util/serializer"
)

type CommentService struct {
	Content string `form:"content"json:"content"` //评论内容
}

func (service *CommentService) Create(uid uint, vid string) serializer.Response {
	var user model.User
	model.DB.First(&user, uid)
	var video model.Video
	model.DB.First(&video, vid)
	code := e.SUCCESS
	comment := model.Comment{
		Uid:     user.ID,
		Video:   video,
		Vid:     video.ID,
		Content: service.Content,
		User:    user,
	}
	video.CommentNum++
	model.DB.Save(video)
	err := model.DB.Create(&comment).Error
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
func (service *CommentService) List(vid string) serializer.Response {
	var comments []model.Comment
	var user model.User
	count := 0
	code := e.SUCCESS
	err := model.DB.Model(&model.Comment{}).Where("vid=?", vid).Count(&count).Find(&comments).Error
	for i := 0; i < count; i++ {
		if comments[i].ReplyTime != 0 {
			var replies []model.Reply
			err1 := model.DB.Model(&model.Reply{}).Where("reply_type=?", 0).Where("reply_id=?", comments[i].ID).Find(&replies).Error
			if err1 != nil {
				code = e.ErrorDatabase
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
					Error:  err1.Error(),
				}
			}
			replies = Find(replies)
			replies1 := serializer.BuildReplies(replies)
			comments[i].Reply = replies1
			model.DB.First(&user, comments[i].Uid)
			comments[i].User = user
		}
	}
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(count))
}
func Find(replies []model.Reply) []model.Reply {
	var user model.User
	for i := 0; i < len(replies); i++ {
		if replies[i].ReplyTime != 0 {
			var replies1 []model.Reply
			model.DB.Model(&model.Reply{}).Where("reply_type=?", 1).Where("reply_id=?", replies[i].ID).Find(&replies1)
			model.DB.First(&user, replies1[i].Uid)
			replies1[i].User = user
			replies1 = Find(replies1)
			replies2 := serializer.BuildReplies(replies1)
			replies[i].Reply = replies2
			model.DB.First(&user, replies[i].Uid)
			replies[i].User = user
		}
	}
	return replies
}
