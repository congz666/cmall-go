package model

import (
	"cmall/cache"
	"strconv"

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

// View 获取点击数
func (product *Products) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (product *Products) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
