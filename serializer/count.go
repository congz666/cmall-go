//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:14:38
 * @LastEditors: congz
 * @LastEditTime: 2020-08-13 09:57:18
 */
package serializer

// Count 数量序列化器
type Count struct {
	FavoriteTotal int `json:"favorite_total"`
	NotPayTotal   int `json:"not_pay_total"`
	PayTotal      int `json:"pay_total"`
}

// BuildCount 序列化轮播图
func BuildCount(favoriteTotal, notPayTotal, payTotal int) Count {
	return Count{
		FavoriteTotal: favoriteTotal,
		NotPayTotal:   notPayTotal,
		PayTotal:      payTotal,
	}
}
