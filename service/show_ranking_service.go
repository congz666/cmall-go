package service

import (
	"cmall/cache"
	"cmall/model"
	"cmall/serializer"
	"fmt"

	"strings"
)

// ShowRankingService 展示排行的服务
type ShowRankingService struct {
}

// Show 获取排行
func (service *ShowRankingService) Show() serializer.Response {
	var products []model.Products

	// 从redis读取点击前十的视频
	pros, _ := cache.RedisClient.ZRevRange(cache.RankKey, 0, 9).Result()

	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(pros, ","))
		err := model.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildProducts(products),
	}
}
