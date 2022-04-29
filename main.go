package main

import (
	"github.com/spf13/viper"
	"go-oj/config"
	"go-oj/controller"
	"go-oj/dao/mysql"
	"go-oj/dao/redis"
	"go-oj/logger"
	"go-oj/router"
	"go-oj/utils/snowflake"
	"go.uber.org/zap"
	"log"
)

func main() {
	// 初始化配置文件
	if err := config.Init(); err != nil {
		log.Println("viper init failed!")
		return
	}

	// 初始化日志库
	if err := logger.Init(); err != nil {
		log.Println("zap init failed!")
		return
	}
	defer zap.L().Sync() // 日志同步到文件系统
	zap.L().Debug("zap log init successfully！")

	// 初始化数据库
	if err := mysql.Init(); err != nil {
		zap.L().Error("mysql init failed. err = ", zap.Error(err))
		return
	}
	defer mysql.Close() // 关闭数据库连接

	if err := redis.Init(); err != nil {
		zap.L().Error("redis init failed., err = ", zap.Error(err))
		return
	}
	defer redis.Close() // 关闭redis连接

	// 初始化雪花算法
	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		log.Fatal(err)
		//zap.L().Error("snowflake init failed, err:", zap.Error(err))
		return
	}

	// 初始化校验器 validator
	if err := controller.InitTrans("zh"); err != nil {
		zap.L().Error("validator init failed ", zap.Error(err))
		return
	}

	// 5 注册路由
	r := router.SetUp()
	_ = r.Run(":8080")
}
