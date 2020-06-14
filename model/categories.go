package model

import (
	"github.com/jinzhu/gorm"
)

// Categories 分类模型
type Categories struct {
	gorm.Model
	CategoryID   uint
	CategoryName string
}
