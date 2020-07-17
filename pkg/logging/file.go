//Package logging ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-17 17:33:44
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:35:46
 */
package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	// LogSavePath 日志存放路径
	LogSavePath = "runtime/logs/"
	// LogSaveName 日志存放名称
	LogSaveName = "log"
	// LogFileExt a
	LogFileExt = "log"
	// TimeFormat 日期格式化
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
