package model

import (
	"github.com/jinzhu/gorm"
)

// Pictures 商品图片模型
type Pictures struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
