//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 15:01:44
 * @LastEditors: congz
 * @LastEditTime: 2020-07-19 14:58:56
 */
package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Admin 管理员模型
type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}

// SetPassword 设置密码
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	admin.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL 封面地址
func (admin *Admin) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(admin.Avatar, oss.HTTPGet, 24*60*60)
	//if strings.Contains(signedGetURL, "http://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
	//signedGetURL := "https://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	//}
	return signedGetURL
}
