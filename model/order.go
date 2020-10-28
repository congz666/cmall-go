//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 11:46:28
 * @LastEditors: congz
 * @LastEditTime: 2020-10-28 12:10:53
 */
package model

import (
	"cmall/cache"
	"cmall/pkg/logging"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Order 订单模型
type Order struct {
	gorm.Model
	UserID       uint
	ProductID    uint
	Num          uint
	OrderNum     uint64
	AddressName  string
	AddressPhone string
	Address      string
	Type         uint
}

//ListenOrder 监听订单是否过期
func ListenOrder() {
	go func() {
		for {
			opt := redis.ZRangeBy{
				Min:    strconv.Itoa(0),
				Max:    strconv.Itoa(int(time.Now().Unix())),
				Offset: 0,
				Count:  10,
			}
			orderList, err := cache.RedisClient.ZRangeByScore(os.Getenv("REDIS_ZSET_KEY"), opt).Result()
			if err != nil {
				logging.Info("redis err:", err)
			}
			if len(orderList) != 0 {
				var numList []int
				for _, v := range orderList {
					i, err := strconv.Atoi(v)
					if err != nil {
						logging.Info("Atoi err:", err)
					}
					numList = append(numList, i)
				}
				if err := DB.Delete(&Order{}, "order_num IN (?)", numList).Error; err != nil {
					logging.Info("myql err:", err)
				}
				if err := cache.RedisClient.ZRem(os.Getenv("REDIS_ZSET_KEY"), orderList).Err(); err != nil {
					logging.Info("redis err:", err)
				}
			}
		}
	}()
}
