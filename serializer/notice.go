//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:14:38
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:02:11
 */
package serializer

import "cmall/model"

// Notice 公告序列化器
type Notice struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"created_at"`
}

// BuildNotice 序列化轮播图
func BuildNotice(item model.Notice) Notice {
	return Notice{
		ID:        item.ID,
		Text:      item.Text,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
