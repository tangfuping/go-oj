package service

import (
	"go-oj/dao/mysql"
	"go-oj/models"
	"go.uber.org/zap"
)

//import (
//	"github.com/gin-gonic/gin"
//	"go-oj/models"
//	"net/http"
//)
//
//// GetProblemList
//// @Tags 公共方法
//// @Summary 问题列表
//// @Param page query int false "请输入当前页，默认第一页"
//// @Param size query int false "size"
//// @Success 200 {string} json "{"code":"200","msg","","data":""}
//// @Router /problem-list [get]
//func GetProblemList(c *gin.Context) {
//	models.GetProblemList()
//	c.String(http.StatusOK, "Get Problem List")
//}

// 根据问题id查找问题
func GetProblemByPid(pid string) (*models.ProblemParams, error) {
	dbModel, err := mysql.FindProblemByPid(pid)
	if err != nil {
		zap.L().Error("获取问题失败")
		return nil, err
	}

	var problemParam models.ProblemParams
	problemParam.ProblemId = dbModel.ProblemId.String

	// 题目标题
	if dbModel.Title.Valid {
		problemParam.Title = dbModel.Title.String
	} else {
		problemParam.Title = "题目暂无标题"
	}

	// 作者信息
	if dbModel.Author.Valid {
		problemParam.Author = dbModel.Author.String
	} else {
		problemParam.Title = "作者信息不详"
	}

	problemParam.Type = dbModel.Type
	problemParam.TimeLimit = dbModel.TimeLimit
	problemParam.MemoryLimit = dbModel.MemoryLimit
	problemParam.StackLimit = dbModel.StackLimit

	// 题目描述
	if dbModel.Description.Valid {
		problemParam.Description = dbModel.Description.String
	} else {
		problemParam.Description = "暂无题目描述信息..."
	}

	// 输入描述
	if dbModel.Input.Valid {
		problemParam.Input = dbModel.Input.String
	} else {
		problemParam.Input = "暂无输入描述..."
	}

	// 输出描述
	if dbModel.Output.Valid {
		problemParam.Output = dbModel.Output.String
	} else {
		problemParam.Output = "暂无输出描述..."
	}

	if dbModel.Example.Valid {
		problemParam.Example = dbModel.Example.String
	} else {
		problemParam.Example = "暂无输入用例..."
	}

	problemParam.Difficulty = dbModel.Difficulty

	if dbModel.Hint.Valid {
		problemParam.Hint = dbModel.Hint.String
	} else {
		problemParam.Hint = "暂无提示信息..."
	}

	if dbModel.SpjLanguage.Valid {
		problemParam.SpjLanguage = dbModel.SpjLanguage.String
	} else {
		problemParam.SpjLanguage = "语言..."
	}

	return &problemParam, nil
}
