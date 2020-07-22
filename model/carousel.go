//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:03:57
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:55:20
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Carousel 轮播图模型
type Carousel struct {
	gorm.Model
	ImgPath string
}
