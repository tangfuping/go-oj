/**
  @author: ZYL
  @date:
  @note
*/
package models

import (
	"time"
)

type ProblemParams struct {
	Id               int        `json:"id"`
	JudgeMode        string     `json:"judge_mode"`
	ProblemId        string     `json:"problem_id"` // 问题编号
	Title            string     `json:"title"`
	Author           string     `json:"author"`
	Type             int        `json:"type"`         // type	int		题目类型 0为ACM,1为OI
	TimeLimit        int        `json:"time_limit"`   // 时间限制
	MemoryLimit      int        `json:"memory_limit"` // 空间限制
	StackLimit       int        `json:"stack_limit"`  // 栈大小限制
	Description      string     `json:"description"`  // 问题描述
	Input            string     `json:"input"`        // 输入描述
	Output           string     `json:"output"`       // 输出描述
	Example          string     `json:"example"`      // 输入样例
	Source           int        `json:"source"`       // 问题来源
	Difficulty       int        `json:"difficulty"`   // 难度
	Hint             string     `json:"hint"`         // 提示
	Auth             int        `json:"auth"`
	IoScore          int        `json:"io_score"` // 当该题目为io题目时的分数 默认为100
	CodeShare        bool       `json:"code_share"`
	SpjCode          string     `json:"spj_code"`
	SpjLanguage      string     `json:"spj_language"`
	UserExtraFile    string     `json:"user_extra_file"`
	IsRemoveEndBlank bool       `json:"is_remove_end_blank"`
	OpenCaseResult   bool       `json:"open_case_result"` // 是否默认开启该题目的测试样例结果查看
	CaseVersion      string     `json:"case_version"`     // 题目测试数据的版本号
	IsUploadCase     bool       `json:"is_upload_case"`   // 是否是上传zip评测数据的
	ModifiedUser     string     `json:"modified_user"`    // 最新修改题目的用户
	GmtCreate        *time.Time `json:"gmt_create"`       // 创建时间
}
