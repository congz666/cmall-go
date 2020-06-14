package model

import (
	"github.com/jinzhu/gorm"
)

// Carousels 轮播图模型
type Carousels struct {
	gorm.Model
	ImgPath string
}
