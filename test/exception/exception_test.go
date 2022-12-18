package exception

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/vansenic/kit/exception"
	"github.com/vansenic/kit/validate"
	"log"
	"os"
	"testing"
)

func init() {
	err := validate.InitValidate(context.Background(), validate.VALIDATOR, validate.Option{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestZapBadValidate(t *testing.T) {
	// 错误样本丨缺失必填项丨要求无法通过验证
	hostname, _ := os.Hostname()
	err := exception.InitLogger(context.Background(), exception.ZAP, exception.Option{
		App:   "kit",
		Label: hostname,
	})
	assert.ErrorContains(t, err, "必填项")
}

func TestZapGood(t *testing.T) {
	// 正确样本丨通过验证
	hostname, _ := os.Hostname()
	err := exception.InitLogger(context.Background(), exception.ZAP, exception.Option{
		Env:   "test",
		App:   "kit",
		Label: hostname,
	})
	if err != nil {
		t.Error(err)
	}
}
