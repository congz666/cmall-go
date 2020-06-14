package model

import (
	"github.com/jinzhu/gorm"
)

// Products 商品模型
type Products struct {
	gorm.Model
	Name          string
	CategoryID    int
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
}
