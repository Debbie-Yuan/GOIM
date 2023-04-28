package main

import (
	"im/router"
	"im/utils"
)

func main() {
	// 初始化系统，加载配置
	utils.InitConfig()
	utils.InitDatabase()
	// 加载路由
	r := router.Router()
	// 启动程序
	r.Run()
}
