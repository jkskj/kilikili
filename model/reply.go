package model

import "github.com/jinzhu/gorm"

type Reply struct {
	gorm.Model
	User      User        `gorm:"ForeignKey:Uid"`
	Uid       uint        `gorm:"not null"` //回复的用户id
	ReplyType int64       `gorm:"not null"` //判断是对评论的回复还是对回复的回复
	ReplyId   uint        `gorm:"not null"` //回复评论的id
	Video     Video       `gorm:"ForeignKey:Vid"`
	Vid       uint        `gorm:"not null"` //视频id
	Content   string      `gorm:"not null"` //评论内容
	Reply     interface{} `gorm:"-"`
	ReplyTime int64       `gorm:"default:0"` //回复数
}
