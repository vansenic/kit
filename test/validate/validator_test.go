package validate

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/vansenic/kit/validate"
	"testing"
)

type Sample struct {
	Name string `json:"name" label:"姓名" validate:"required"`
	Age  int    `json:"age" label:"年龄" validate:"gt=18"`
	Link string `json:"link" label:"个人主页" validate:"required,url,startswith=http"`
}

func TestValidator(t *testing.T) {
	err := validate.InitValidate(context.Background(), validate.VALIDATOR, validate.Option{})
	if err != nil {
		t.Error(err.Error())
	}
	// 错误样本丨缺失必填项丨要求无法通过验证
	sampleBadMiss := Sample{Name: "", Age: 25, Link: "https://github.com/vansenic"}
	message, _ := validate.Validate.Work(sampleBadMiss)
	assert.Equal(t, message, "姓名为必填项")
	// 错误样本丨格式不正确丨要求无法通过验证
	sampleBadType := Sample{Name: "张良", Age: 25, Link: "github.com/vansenic"}
	message, _ = validate.Validate.Work(sampleBadType)
	assert.Equal(t, message, "个人主页的值必须符合网址格式")
	// 错误样本丨值超限丨要求无法通过验证
	sampleBadLimit := Sample{Name: "张良", Age: 16, Link: "https://github.com/vansenic"}
	message, _ = validate.Validate.Work(sampleBadLimit)
	assert.Equal(t, message, "年龄的值必须大于18")
	// 错误样本丨值超限丨要求无法通过验证
	sampleBadStartWith := Sample{Name: "张良", Age: 25, Link: "ftp://github.com/vansenic"}
	message, _ = validate.Validate.Work(sampleBadStartWith)
	assert.Equal(t, message, "个人主页的值必须以[http]为开头")
	// 正确样本丨值超限丨要求无法通过验证
	sampleGood := Sample{Name: "张良", Age: 25, Link: "https://github.com/vansenic"}
	message, _ = validate.Validate.Work(sampleGood)
	assert.Equal(t, message, "")
}
