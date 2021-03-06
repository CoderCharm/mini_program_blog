/*
* @Time    : 2020-11-17 11:47
* @Author  : CoderCharm
* @File    : server.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 启动服务
**/
package core

import (
	"mini_blog_server/initialize"
)

func RunWindowsServer(addr string) {

	// 初始化redis
	initialize.Redis()

	// 初始化路由
	Router := initialize.Routers()

	_ = Router.Run(addr)
}
