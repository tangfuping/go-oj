package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-oj/docs"
	"go-oj/logger"
	"go-oj/service"
	"net/http"
)

func SetUp() *gin.Engine {
	r := gin.New() // 创建一个路由
	mode := viper.GetString("app.mode")
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 整合zap日志

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	// Swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ping", service.Ping)
	//r.GET("/problem-list", service.GetProblemList)

	return r

}

func Router() *gin.Engine {
	// 1.创建路由
	r := gin.Default()

	// Swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ping", service.Ping)
	//r.GET("/problem-list", service.GetProblemList)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")

	return r
}
