package main

import (
	"gin-gorm-oj/dao/mysql"
	"gin-gorm-oj/router"
	"go.uber.org/zap"
)

func main() {
	r := router.Router()
	r.Run(":8080")

	if err := mysql.Init(); err != nil {
		zap.L().Error("mysql init fail!")
	}

}
