package config

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/vansenic/kit/config"
	"github.com/vansenic/kit/validate"
	"log"
	"testing"
)

func init() {
	err := validate.InitValidate(context.Background(), validate.VALIDATOR, validate.Option{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestViperBadPath(t *testing.T) {
	// 错误样本丨路径不正确丨要求无法通过验证
	err := config.InitConfig(context.Background(), config.VIPER, config.OptionConfig{
		Path:     "/",
		Filename: "config.yml",
		Type:     config.YML,
	})
	assert.ErrorContains(t, err, "Not Found")
}

func TestViperBadType(t *testing.T) {
	// 错误样本丨类型不正确丨要求无法通过验证
	err := config.InitConfig(context.Background(), config.VIPER, config.OptionConfig{
		Path:     "./",
		Filename: "default_config.yml",
		Type:     config.JSON,
	})
	config.Con.GetString("name")
	assert.ErrorContains(t, err, "invalid")
}

func TestViperGood(t *testing.T) {
	// 正确样本丨通过验证
	err := config.InitConfig(context.Background(), config.VIPER, config.OptionConfig{
		Path:     "./",
		Filename: "default_config.yml",
		Type:     config.YML,
	})
	name := config.Con.GetString("name")
	if err != nil {
		t.Error(err)
	}
	if name != "kit" {
		t.Error("name 的值不正确")
	}
}
