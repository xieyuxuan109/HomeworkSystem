package main

import (
	"log"

	"github.com/xieyuxuan109/homeworksystem/configs"
	"github.com/xieyuxuan109/homeworksystem/router"
)

func main() {
	err := configs.InitDB()
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	r := router.SetupRouter()
	log.Println("服务器启动在 :8080")
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("启动失败:", err)
	}
}
