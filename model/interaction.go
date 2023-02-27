package model

import "github.com/jinzhu/gorm"

type Interaction struct {
	gorm.Model
	User   User  `gorm:"ForeignKey:Uid"`
	Uid    uint  `gorm:"not null"` //用户id
	Video  Video `gorm:"ForeignKey:Vid"`
	Vid    uint  `gorm:"not null"`  //视频id
	Type   int64 `gorm:"not null"`  //互动类型
	Status int64 `gorm:"default:1"` //状态
}
