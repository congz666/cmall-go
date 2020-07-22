//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 08:59:06
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:54:26
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Favorite 收藏夹模型
type Favorite struct {
	gorm.Model
	UserID    uint
	ProductID uint
}
