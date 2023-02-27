package serializer

import "kilikili/model"

type Reply struct {
	ID        uint        `json:"id" example:"1"`  // 评论id
	Uid       uint        `json:"uid" example:"1"` //用户id
	Vid       uint        `json:"vid" example:"1"` //视频id
	Content   string      `json:"content"`         //弹幕内容
	Reply     interface{} `json:"reply"`
	ReplyTime int64       `json:"reply_time"` //回复数
	ReplyType int64       `json:"reply_type"` //判断是对评论的回复还是对回复的回复
	ReplyId   uint        `json:"reply_id"`   //回复评论的id
	User      User        `json:"user"`
}

func BuildReply(item model.Reply) Reply {
	return Reply{
		ID:        item.ID,
		Uid:       item.Uid,
		Vid:       item.Vid,
		Content:   item.Content,
		Reply:     item.Reply,
		ReplyTime: item.ReplyTime,
		ReplyId:   item.ReplyId,
		ReplyType: item.ReplyType,
		User:      BuildUser(item.User),
	}
}
func BuildReplies(items []model.Reply) (Replies []Reply) {
	for _, item := range items {
		Reply := BuildReply(item)
		Replies = append(Replies, Reply)
	}
	return Replies
}
