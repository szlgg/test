package main

import (
	"fmt"
	"test/app/system/index"
	"test/config"
)

func main() {
	// app.Run()

	// 加载配置文件
	if err := config.Init("config/config.ini"); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 业务系统初始化
	router := index.Init()

	// 静态目录设置
	router.Static("/public", "./public")
	router.Static("/upload", "./upload")

	err := router.Run(config.Conf.Server.Address)
	if err != nil {
		return
	}
}
