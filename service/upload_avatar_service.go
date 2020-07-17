//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:58:02
 */
package service

import (
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
	"mime"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// UploadAvatarService 获得上传oss token的服务
type UploadAvatarService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadAvatarService) Post() serializer.Response {
	code := e.SUCCESS
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		logging.Info(err)
		code = e.ERROR_OSS
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		logging.Info(err)
		code = e.ERROR_OSS
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取扩展名
	ext := filepath.Ext(service.Filename)

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_OSS
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_OSS
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
