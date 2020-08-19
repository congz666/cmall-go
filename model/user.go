// Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-18 19:33:07
 */
package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordDigest string
	Nickname       string `gorm:"unique"`
	Status         string
	Limit          int
	Avatar         string `gorm:"size:1000"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL 封面地址
func (user *User) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(user.Avatar, oss.HTTPGet, 24*60*60)
	//if strings.Contains(signedGetURL, "http://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
	//signedGetURL := "https://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	//}
	return signedGetURL
}
