//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-21 23:05:55
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:08:52
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// ProductInfoImg 商品图片模型
type ProductInfoImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
