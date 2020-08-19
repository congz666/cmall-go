// Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-18 19:37:20
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// UserAuth 用户权限模型
type UserAuth struct {
	gorm.Model
	UserID       uint
	IdentityType string //第三方应用名称 (微信 , 微博等)
	Identifier   string `gorm:"unique"` //标识 (第三方应用的唯一标识)
	Token        string //token凭证 (保存 token)
	RefreshToken string
}
