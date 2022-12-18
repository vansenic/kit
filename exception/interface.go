package exception

import (
	"context"
	"errors"
	"github.com/vansenic/kit/code"
	"github.com/vansenic/kit/rank"
	"github.com/vansenic/kit/validate"
	"sync"
)

type SupLogger string

const ZAP SupLogger = "zap"

var Log Logger
var once sync.Once

type Logger interface {
	Debug(ctx context.Context, message Message)
	Info(ctx context.Context, message Message)
	Warn(ctx context.Context, message Message)
	Error(ctx context.Context, message Message)
	Fatal(ctx context.Context, message Message)
	Shutdown(ctx context.Context)
}

type Option struct {
	Env   string `json:"env" label:"当前环境" validate:"required"`
	App   string `json:"app" label:"应用名称" validate:"required"`
	Label string `json:"label" label:"标记"`
}

type Message struct {
	Rank      rank.Rank `json:"rank" label:"数据级别"`
	Code      code.Code `json:"code" label:"参照码"`
	Final     bool      `json:"final" label:"出口日志"`
	Traceback string    `json:"traceback" label:"错误栈信息"`
	Hidden    bool      `json:"hidden " label:"是否隐藏"`
	Remark    string    `json:"remark" label:"备注信息"`
	Inp       string    `json:"inp" label:"输入"`
	Oup       string    `json:"oup" label:"输出"`
	Line      string    `json:"line" label:"位置丨行"`
}

func InitLogger(ctx context.Context, sup SupLogger, option Option) error {
	var err error
	var temp = Log
	// validate
	if message, err := validate.Validate.Work(option); err != nil {
		return errors.New(message)
	}
	// create
	once.Do(func() {
		if temp, err = FactoryLogger(ctx, sup, option); err == nil {
			if temp == nil {
				err = errors.New("初始化失败")
			}
		}
		Log = temp
	})
	return err
}

func FactoryLogger(ctx context.Context, sup SupLogger, option Option) (Logger, error) {
	switch sup {
	case ZAP:
		return NewLoggerZap(ctx, option)
	default:
		return NewLoggerZap(ctx, option)
	}
}
