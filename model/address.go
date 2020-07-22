//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 10:38:13
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:55:33
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Address 收货地址模型
type Address struct {
	gorm.Model
	UserID  uint
	Name    string
	Phone   string
	Address string
}
