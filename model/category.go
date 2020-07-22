//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 22:09:35
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:54:47
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
}
