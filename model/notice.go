//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-08-04 10:50:29
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 10:52:42
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Notice 公告模型
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
