/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-10-28 11:58:57
 */
package main

import (
	"cmall/conf"
	"cmall/model"
	"cmall/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	model.ListenOrder()
	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
