package service

import (
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
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "打开文件失败",
			Error:  err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "读取文件失败",
			Error:  err.Error(),
		}
	}
	str := string(bytes)
	return serializer.Response{
		Data: str,
	}
}
