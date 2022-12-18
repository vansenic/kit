package exception

import (
	"context"
	"github.com/vansenic/kit/code"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type LoggerZap struct {
	core  *zap.Logger
	Env   string `json:"env"`
	App   string `json:"app"`
	Label string `json:"label"`
}

func NewLoggerZap(ctx context.Context, option OptionLogger) (Logger, error) {
	enc := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "traceback",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	// 构建输出方式
	zapWriterSync := zapcore.AddSync(os.Stdout)
	// 构建输出格式
	encoder := zapcore.NewJSONEncoder(enc)
	// 应用参数
	zapCore := zapcore.NewCore(
		encoder,
		zapWriterSync,
		zapcore.InfoLevel,
	)
	logger := zap.New(zapCore)
	return &LoggerZap{
		Env:   option.Env,
		App:   option.App,
		Label: option.Label,
		core:  logger,
	}, nil
}

func (c *LoggerZap) write(ctx context.Context, level string, message Message) {
	text := code.GetCommentaries(message.Code)
	if message.Code == 0 {
		message.Code = code.CodeDefault
	}
	if message.Rank == "" {
		message.Remark = "info"
	}
	// 将传入的参数应用到日志
	content := []zapcore.Field{
		zap.String("trace", ctx.Value("trace").(string)),
		zap.String("rank", string(message.Rank)),
		zap.String("env", c.Env),
		zap.String("app", c.App),
		zap.String("label", c.Label),
		zap.String("traceback", message.Traceback),
		zap.Int("code", int(message.Code)),
		zap.Bool("final", message.Final),
		zap.Bool("hidden", false),
		zap.String("inp", message.Inp),
		zap.String("oup", message.Oup),
		zap.String("remark", message.Remark),
		zap.String("line", message.Line),
	}
	// 调用库原生方法写日志
	switch level {
	case "debug":
		c.core.Debug(text, content...)
	case "warn":
		c.core.Warn(text, content...)
	case "error":
		c.core.Error(text, content...)
	case "fatal":
		c.core.Fatal(text, content...)
	default:
		c.core.Info(text, content...)
	}
}

func (c *LoggerZap) Debug(ctx context.Context, message Message) {
	c.write(ctx, "debug", message)
}

func (c *LoggerZap) Info(ctx context.Context, message Message) {
	c.write(ctx, "info", message)
}

func (c *LoggerZap) Warn(ctx context.Context, message Message) {
	c.write(ctx, "warn", message)
}

func (c *LoggerZap) Error(ctx context.Context, message Message) {
	c.write(ctx, "error", message)
}

func (c *LoggerZap) Fatal(ctx context.Context, message Message) {
	c.write(ctx, "fatal", message)
}

func (c *LoggerZap) Shutdown(ctx context.Context) {

}
