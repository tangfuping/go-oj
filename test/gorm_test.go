package test

import (
	"database/sql"
	"go-oj/dao/mysql"
	"go-oj/models"
	"testing"
	"time"
)

func TestGormTest(t *testing.T) {
	//// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:123@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//data := make([]*models.Problem, 0)
	//err = db.Find(&data).Error
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for _, v := range data {
	//	fmt.Printf("Problem ===> %v \n", v)
	//
	//}
}

// TestInsertProblems 测试插入问题数据
func TestInsertProblems(t *testing.T) {
	curTime := time.Now()
	var model = models.ProblemModel{
		Id:               1,
		JudgeMode:        sql.NullString{String: "default", Valid: true},
		ProblemId:        sql.NullString{String: "1001", Valid: true},
		Title:            sql.NullString{String: "两数之和", Valid: true},
		Author:           sql.NullString{String: "郑裕龙", Valid: true},
		Type:             1,
		TimeLimit:        2000,
		MemoryLimit:      2,
		StackLimit:       2,
		Description:      sql.NullString{String: "给定一组无序数组nums和一个目标值trg，在nums中找出两个值x,y，使得x和y的和为，答案不唯一！", Valid: true},
		Input:            sql.NullString{String: "nums=[1,3,5,4,8,6], trg=12", Valid: true},
		Output:           sql.NullString{String: "4 8", Valid: true},
		Example:          sql.NullString{String: "", Valid: true},
		Source:           2,
		Difficulty:       0,
		Hint:             sql.NullString{String: "", Valid: true},
		Auth:             1,
		IoScore:          100,
		CodeShare:        true,
		SpjCode:          sql.NullString{String: "", Valid: true},
		SpjLanguage:      sql.NullString{String: "Golang C++ Java", Valid: true},
		UserExtraFile:    sql.NullString{String: "", Valid: true},
		IsRemoveEndBlank: false,
		OpenCaseResult:   false,
		CaseVersion:      sql.NullString{String: "v1.0", Valid: true},
		IsUploadCase:     true,
		ModifiedUser:     sql.NullString{String: "123456", Valid: true},
		GmtCreate:        &curTime,
		GmtModified:      &curTime,
	}

	_ = mysql.InsertIntoProblem(&model)
}
