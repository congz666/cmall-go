package model

import (
	"github.com/jinzhu/gorm"
)

// Favorites 收藏夹模型
type Favorites struct {
	gorm.Model
	UserID    uint
	ProductID uint
}
