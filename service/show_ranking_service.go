/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-02 11:52:43
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:50:52
 */
package service

import (
	"cmall/cache"
	"cmall/model"
	"cmall/pkg/e"
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
	code := e.SUCCESS
	// 从redis读取点击前十的视频
	pros, _ := cache.RedisClient.ZRevRange(cache.RankKey, 0, 9).Result()

	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(pros, ","))
		err := model.DB.Where("id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			code := e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
