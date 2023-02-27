package serializer

import "kilikili/model"

type Comment struct {
	ID        uint        `json:"id" example:"1"`  // 评论id
	Uid       uint        `json:"uid" example:"1"` //用户id
	Vid       uint        `json:"vid" example:"1"` //视频id
	Content   string      `json:"content"`         //弹幕内容
	Reply     interface{} `json:"reply"`
	ReplyTime int64       `json:"reply_time"` //回复数
	User      User        `json:"user"`
}

func BuildComment(item model.Comment) Comment {
	return Comment{
		ID:        item.ID,
		Uid:       item.Uid,
		Vid:       item.Vid,
		Content:   item.Content,
		Reply:     item.Reply,
		ReplyTime: item.ReplyTime,
		User:      BuildUser(item.User),
	}
}
func BuildComments(items []model.Comment) (Comments []Comment) {
	for _, item := range items {
		Comment := BuildComment(item)
		Comments = append(Comments, Comment)
	}
	return Comments
}
