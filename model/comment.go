package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	User      User        `gorm:"ForeignKey:Uid"`
	Uid       uint        `gorm:"not null"` //用户id
	Video     Video       `gorm:"ForeignKey:Vid"`
	Vid       uint        `gorm:"not null"` //视频id
	Content   string      `gorm:"not null"` //评论内容
	Reply     interface{} `gorm:"-"`
	ReplyTime int64       `gorm:"default:0"` //回复数
}
