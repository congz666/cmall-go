/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 23:12:42
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:44:18
 */
package service

import (
	"cmall/pkg/e"
	"cmall/serializer"
	"io/ioutil"
	"os"
)

// ReadMeService 发送README的服务
type ReadMeService struct {
}

// Read
func (service *ReadMeService) Read() serializer.Response {
	f, err := os.Open("./me.md")
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	str := string(bytes)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   str,
	}
}
