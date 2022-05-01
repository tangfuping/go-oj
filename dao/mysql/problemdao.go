/**
  @author: ZYL
  @date:
  @note
*/
package mysql

import (
	"fmt"
	"go-oj/models"
	"go.uber.org/zap"
)

// 根据问题id查询数据
func FindProblemByPid(problemId string) (*models.ProblemModel, error) {
	sqlStr := `SELECT * FROM tb_problem WHERE problem_id = ?`
	var model models.ProblemModel
	if err := db.Get(&model, sqlStr, problemId); err != nil {
		zap.L().Error("查询问题失败，err=", zap.Error(err))
		return nil, err
	}

	return &model, nil
}

// 插入问题数据

func InsertIntoProblem(model *models.ProblemModel) error {

	sqlStr := `INSERT INTO 
				tb_problem(
				judge_mode,problem_id,title,author,type,
				time_limit,memory_limit,stack_limit,description,
				input,output,example,source,difficulty,hint,auth,
				io_score,code_share,spj_code,spj_language,user_extra_file,
				is_remove_end_blank,open_case_result,case_version,is_upload_case,
				modified_user, gmt_create, gmt_modified) 
				VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	fmt.Println(db)
	_, err := db.Exec(sqlStr,
		model.JudgeMode.String, model.ProblemId.String, model.Title.String, model.Author.String, model.Type,
		model.TimeLimit, model.MemoryLimit, model.StackLimit, model.Description.String,
		model.Input.String, model.Output.String, model.Example.String, model.Source, model.Difficulty, model.Hint.String,
		model.Auth, model.IoScore, model.CodeShare, model.SpjCode.String, model.SpjLanguage.String, model.UserExtraFile.String,
		model.IsRemoveEndBlank, model.OpenCaseResult, model.CaseVersion.String, model.IsUploadCase, model.ModifiedUser.String,
		model.GmtCreate, model.GmtModified)
	if err != nil {
		zap.L().Error("保存问题失败，err=", zap.Error(err))
		return err
	}

	return nil
}
