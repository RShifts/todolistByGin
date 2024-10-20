package main

import (
	"github.com/gin-gonic/gin"
	"todolist/db"
	"todolist/router"
)

func main() {
	var r = gin.Default()
	//var DB *gorm.DB
	// 初始化路由
	router.Init(r)
	// 初始化数据库
	db.Init()
	err := r.Run()
	if err != nil {
		return
	}
}
