//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:08:41
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:03:11
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowParamImgsService 商品参数图片详情的服务
type ShowParamImgsService struct {
}

// Show 商品图片
func (service *ShowParamImgsService) Show(id string) serializer.Response {
	var paramImgs []model.ProductParamImg
	code := e.SUCCESS

	err := model.DB.Where("product_id=?", id).Find(&paramImgs).Error
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
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildParamImgs(paramImgs),
	}
}
