package service

import (
	"kilikili/model"
	"kilikili/util/e"
	"kilikili/util/serializer"
)

type ReplyService struct {
	ReplyType int64  `form:"reply_type" json:"reply_type"` //判断是对评论的回复还是对回复的回复
	ReplyId   uint   `form:"reply_id" json:"reply_id"`     //回复评论的id
	Content   string `form:"content" json:"content"`       //评论内容
}

// Create 上传回复
func (service *ReplyService) Create(uid uint, vid string) serializer.Response {
	var user model.User
	model.DB.First(&user, uid)
	var video model.Video
	model.DB.First(&video, vid)
	code := e.SUCCESS
	reply := model.Reply{
		User:      user,
		Uid:       user.ID,
		Video:     video,
		Vid:       video.ID,
		Content:   service.Content,
		ReplyId:   service.ReplyId,
		ReplyType: service.ReplyType,
	}
	video.CommentNum++
	model.DB.Save(video)
	if service.ReplyType == 1 {
		var reply1 model.Reply
		err1 := model.DB.First(&reply1, service.ReplyId).Error
		if err1 != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err1.Error(),
			}
		}
		reply1.ReplyTime = reply1.ReplyTime + 1
		model.DB.Save(reply1)
	} else {
		var comment model.Comment
		err1 := model.DB.First(&comment, service.ReplyId).Error
		if err1 != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err1.Error(),
			}
		}
		comment.ReplyTime = comment.ReplyTime + 1
		model.DB.Save(comment)
	}
	err := model.DB.Create(&reply).Error
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
