package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func Init(r *gin.Engine) {
	// 日志
	createLog()
	// 路由
	Router(r)
}

func createLog() {
	gin.DisableConsoleColor()
	var logName = fmt.Sprintf("gin_%s.log", time.Now().Format("2024-10-29"))
	create, _ := os.Create(logName)
	gin.DefaultWriter = io.MultiWriter(create)
}
