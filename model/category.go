//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 22:09:35
 * @LastEditors: congz
 * @LastEditTime: 2020-09-24 13:39:01
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Category 分类模型
type Category struct {
	gorm.Model
	CategoryID   uint
	CategoryName string
	Num          uint
}
