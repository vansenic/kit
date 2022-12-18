package config

import (
	"context"
	"github.com/spf13/viper"
	"time"
)

type Viper struct {
	core *viper.Viper
}

func NewConfigViper(ctx context.Context, option OptionConfig) (InterfaceConfig, error) {
	vip := viper.New()
	vip.AddConfigPath(option.Path)
	vip.SetConfigType(string(option.Type))
	vip.SetConfigName(option.Filename)
	err := vip.ReadInConfig()
	return &Viper{core: vip}, err
}

func (c *Viper) Get(key string) interface{} {
	return c.core.Get(key)
}

func (c *Viper) GetBool(key string) bool {
	return c.core.GetBool(key)
}

func (c *Viper) GetString(key string) string {
	return c.core.GetString(key)
}

func (c *Viper) GetInt(key string) int {
	return c.core.GetInt(key)
}

func (c *Viper) GetInt64(key string) int64 {
	return c.core.GetInt64(key)
}

func (c *Viper) GetDuration(key string) time.Duration {
	return c.core.GetDuration(key)
}

func (c *Viper) GetStringSlice(key string) []string {
	return c.core.GetStringSlice(key)
}

func (c *Viper) GetIntSlice(key string) []int {
	return c.core.GetIntSlice(key)
}

func (c *Viper) GetStringMap(key string) map[string]interface{} {
	return c.core.GetStringMap(key)
}

func (c *Viper) GetStringMapString(key string) map[string]string {
	return c.core.GetStringMapString(key)
}

func (c *Viper) GetStringMapStringSlice(key string) map[string][]string {
	return c.core.GetStringMapStringSlice(key)
}
