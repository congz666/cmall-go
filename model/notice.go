//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-08-04 10:50:29
 * @LastEditors: congz
 * @LastEditTime: 2020-08-09 10:17:59
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Notice 公告模型 存放公告和邮件模板
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
