//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 15:30:56
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:55:10
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Cart 订单模型
type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Num       uint
	MaxNum    uint
	Check     bool
}
