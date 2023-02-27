package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`   //唯一索引
	PasswordDigest string `gorm:"not null"` //密文
	Email          string `gorm:"not null"`
	Avatar         string `gorm:"default:0"` //头像
	Role           int64  `gorm:"default:0"` //权限
	IsBlock        int64  `gorm:"default:0"//是否被封`
}

// SetPassword 加密密码
func (user *User) SetPassword(password string) error {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 验证密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
