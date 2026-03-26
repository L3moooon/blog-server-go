package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"blog_backend_go/config"
	"blog_backend_go/utils"
)

func main() {
	// 1. 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080" // 默认端口
	}

	// 2. 初始化数据库
	config.InitDB()

	// 3. 执行模型自动迁移
	utils.AutoMigrate(config.DB)

	// 4. 创建 Gin 路由
	r := gin.Default()

	// 示例接口
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin & GORM 环境已就绪！",
		})
	})

	// 5. 启动服务
	log.Printf("🚀 服务已启动：http://localhost:%s", port)
	r.Run(":" + port)
}
