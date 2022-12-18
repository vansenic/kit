package validate

import (
	"context"
	"errors"
	"sync"
)

var once sync.Once
var Validate InterfaceValidate

type SupValidate string

var VALIDATOR SupValidate = "validator"

type Option struct {
}

type InterfaceValidate interface {
	Work(payload interface{}) (string, error)
}

func InitValidate(ctx context.Context, sup SupValidate, option Option) error {
	var err error
	var temp = Validate
	// create
	once.Do(func() {
		if temp, err = FactoryValidate(ctx, sup, option); err == nil {
			if temp == nil {
				err = errors.New("初始化失败")
			}
		}
		Validate = temp
	})
	return err
}

func FactoryValidate(ctx context.Context, sup SupValidate, option Option) (InterfaceValidate, error) {
	switch sup {
	case VALIDATOR:
		return NewValidator(ctx, option)
	default:
		return NewValidator(ctx, option)
	}
}
