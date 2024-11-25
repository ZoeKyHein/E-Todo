package main

import (
	"E-Todo/config"
	"E-Todo/routes"
)

func main() {
	// 初始化数据库连接
	config.InitDB()
	r := routes.SetupRouter()
	// 启动服务器
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
