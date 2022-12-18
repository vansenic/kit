package exception

import (
	"github.com/vansenic/kit/code"
	"github.com/vansenic/kit/rank"
)

type Exception struct {
	Rank      rank.Rank `json:"rank" label:"数据级别"`
	Code      code.Code `json:"code" label:"参照码"`
	Traceback string    `json:"traceback" label:"错误栈信息"`
	Remark    string    `json:"remark" label:"备注信息"`
	Error     bool      `json:"error" label:"是否错误"`
	Line      string    `json:"line" label:"位置丨行"`
}
