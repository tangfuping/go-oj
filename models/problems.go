/**
  @author: ZYL
  @date:
  @note
*/
package models

import (
	"database/sql"
	"time"
)

type ProblemModel struct {
	Id               int            `db:"id" json:"id"`
	JudgeMode        sql.NullString `db:"judge_mode"`
	ProblemId        sql.NullString `db:"problem_id"`
	Title            sql.NullString `db:"title"`
	Author           sql.NullString `db:"author"`
	Type             int            `db:"type"` // type	int		题目类型 0为ACM,1为OI
	TimeLimit        int            `db:"time_limit"`
	MemoryLimit      int            `db:"memory_limit"`
	StackLimit       int            `db:"stack_limit"`
	Description      sql.NullString `db:"description"`
	Input            sql.NullString `db:"input"`
	Output           sql.NullString `db:"output"`
	Example          sql.NullString `db:"example"`
	Source           int            `db:"source"`
	Difficulty       int            `db:"difficulty"`
	Hint             sql.NullString `db:"hint"`
	Auth             int            `db:"auth"`
	IoScore          int            `db:"io_score"` // 当该题目为io题目时的分数 默认为100
	CodeShare        bool           `db:"code_share"`
	SpjCode          sql.NullString `db:"spj_code"`
	SpjLanguage      sql.NullString `db:"spj_language"`
	UserExtraFile    sql.NullString `db:"user_extra_file"`
	IsRemoveEndBlank bool           `db:"is_remove_end_blank"`
	OpenCaseResult   bool           `db:"open_case_result"` // 是否默认开启该题目的测试样例结果查看
	CaseVersion      sql.NullString `db:"case_version"`     // 题目测试数据的版本号
	IsUploadCase     bool           `db:"is_upload_case"`   // 是否是上传zip评测数据的
	ModifiedUser     sql.NullString `db:"modified_user"`    // 最新修改题目的用户
	GmtCreate        *time.Time     `db:"gmt_create"`       // 创建时间
	GmtModified      *time.Time     `db:"gmt_modified"`     // 修改时间
}
