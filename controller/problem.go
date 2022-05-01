/**
  @author: ZYL
  @date:
  @note
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"go-oj/service"
	"net/http"
)

func GetProblemByPidHandler(context *gin.Context) {
	pid := context.Param("pid") // 获取手机
	result, err := service.GetProblemByPid(pid)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
