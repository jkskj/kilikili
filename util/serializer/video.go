package serializer

import (
	"kilikili/model"
)

type Video struct {
	ID            uint   `json:"id" ` // 视频ID
	Uid           uint   `json:"uid"`
	Title         string `json:"title" `         // 标题
	Content       string `json:"content" `       // 视频路径
	View          uint64 `json:"view" `          // 浏览量
	CreatedAt     int64  `json:"created_at"`     //创建时间
	Uptime        int64  `json:"uptime"`         //上传时间
	CommentNum    int64  `json:"comment_num"`    //评论数
	CollectionNum int64  `json:"collection_num"` //评论数
	LikeNum       int64  `json:"like_num"`       //收藏数
	ForwardNum    int64  `json:"forward_num"`    //转发数
	User          User   `json:"user"`
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID:            item.ID,
		Uid:           item.Uid,
		Title:         item.Title,
		Content:       item.Content,
		User:          BuildUser(item.User),
		View:          item.View(),
		CreatedAt:     item.CreatedAt.Unix(),
		Uptime:        item.Uptime,
		CommentNum:    item.CommentNum,
		CollectionNum: item.CollectionNum,
		LikeNum:       item.LikeNum,
		ForwardNum:    item.ForwardNum,
	}
}
