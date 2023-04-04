package serializer

import "kilikili/model"

type BulletComment struct {
	ID      uint   `json:"id" example:"1"`  // 弹幕id
	Uid     uint   `json:"uid" example:"1"` //用户id
	Vid     uint   `json:"vid" example:"1"` //视频id
	Content string `json:"content"`         //弹幕内容
	AddTime string `json:"add_time"`        //弹幕出现在视频的时间
}

func BuildBulletComment(item model.BulletComment) BulletComment {
	return BulletComment{
		ID:      item.ID,
		Uid:     item.Uid,
		Vid:     item.Vid,
		Content: item.Content,
		AddTime: item.AddTime,
	}
}
func BuildBulletComments(items []model.BulletComment) (BulletComments []BulletComment) {
	for _, item := range items {
		BulletComment := BuildBulletComment(item)
		BulletComments = append(BulletComments, BulletComment)
	}
	return BulletComments
}
