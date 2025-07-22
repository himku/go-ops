package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 其他路由分组可在此添加
	// api := r.Group("/api")
	// {
	//     api.GET("/example", exampleHandler)
	// }

	return r
}
