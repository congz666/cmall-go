//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:08:41
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:49:10
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowInfoImgsService 商品详情图片详情的服务
type ShowInfoImgsService struct {
}

// Show 商品图片
func (service *ShowInfoImgsService) Show(id string) serializer.Response {
	var infoImgs []model.ProductInfoImg
	code := e.SUCCESS

	err := model.DB.Where("product_id=?", id).Find(&infoImgs).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildInfoImgs(infoImgs),
	}
}
