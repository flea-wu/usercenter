package main

import "github.com/gin-gonic/gin"

func main() {

	// 1. 安装  gin : go get -u github.com/gin-gonic/gin
	// 2. go.mod 中有对应的依赖信息

	// 获取一个路由
	r := gin.Default()

	// get 方法 和 处理get 方法的逻辑
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "tiaozao",
		})
	})

	// 启动
	_ = r.Run()
}
