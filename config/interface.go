package config

import (
	"context"
	"errors"
	"github.com/vansenic/kit/validate"
	"sync"
	"time"
)

type SupConfig string
type SupType string

var once sync.Once
var Con InterfaceConfig

var VIPER SupConfig = "viper"
var YAML SupType = "yaml"
var YML SupType = "yml"
var JSON SupType = "json"

type OptionConfig struct {
	Path     string  `json:"path"`
	Type     SupType `json:"type"`
	Filename string  `json:"filename"`
	Author   bool    `json:"author"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

type InterfaceConfig interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt64(key string) int64
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
	GetIntSlice(key string) []int
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
}

func InitConfig(ctx context.Context, sup SupConfig, option OptionConfig) error {
	var err error
	var temp = Con
	// validate
	if message, err := validate.Validate.Work(option); err != nil {
		return errors.New(message)
	}
	// create
	once.Do(func() {
		if temp, err = FactoryConfig(ctx, sup, option); err == nil {
			if temp == nil {
				err = errors.New("初始化失败")
			}
		}
		Con = temp
	})
	return err
}

func FactoryConfig(ctx context.Context, sup SupConfig, option OptionConfig) (InterfaceConfig, error) {
	switch sup {
	case VIPER:
		return NewConfigViper(ctx, option)
	default:
		return NewConfigViper(ctx, option)
	}
}
