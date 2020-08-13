//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 13:35:13
 * @LastEditors: congz
 * @LastEditTime: 2020-08-13 14:01:45
 */
package service

import (
	"cmall/pkg/e"
	"cmall/serializer"
	"fmt"
	"net/url"
)

// InitQQService 商品详情的服务
type InitQQService struct {
}

// Init QQ初始化
func (service *InitQQService) Init() serializer.Response {
	code := e.SUCCESS
	responseType := "code"
	clientID := "101898836"
	redirectURL := "http://cmall.congz.top/"
	state := "1"
	path := url.QueryEscape(redirectURL)

	loginURL := fmt.Sprintf("https://graph.qq.com/oauth2.0/authorize?response_type=%s&client_id=%s&redirect_url=%s&state=%s", responseType, clientID, path, state)

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   loginURL,
	}
}
