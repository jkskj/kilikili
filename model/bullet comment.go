package model

import "github.com/jinzhu/gorm"

type BulletComment struct {
	gorm.Model
	User    User   `gorm:"ForeignKey:Uid"`
	Uid     uint   `gorm:"not null"` //用户id
	Video   Video  `gorm:"ForeignKey:Vid"`
	Vid     uint   `gorm:"not null"` //视频id
	Content string `gorm:"not null"` //弹幕内容
	AddTime string `gorm:"not null"` //弹幕出现在视频的时间
}
