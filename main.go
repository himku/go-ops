package main

import (
	"log"

	"go-ops/internal/router"

	"github.com/spf13/viper"
)

func main() {
	// 配置初始化
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	r := router.SetupRouter()
	r.Run(":" + port)
}
