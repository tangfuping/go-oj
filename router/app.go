package router

import (
	_ "gin-gorm-oj/docs"
	"gin-gorm-oj/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 1.创建路由
	r := gin.Default()

	// Swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ping", service.Ping)
	r.GET("/problem-list", service.GetProblemList)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")

	return r
}
