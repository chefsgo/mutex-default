package mutex

import (
	"errors"
	"sync"
	"time"

	"github.com/chefsgo/chef"
)

//默认mutex驱动

type (
	defaultMutexDriver  struct{}
	defaultMutexConnect struct {
		name    string
		config  chef.MutexConfig
		setting defaultMutexSetting
		locks   sync.Map
	}
	defaultMutexSetting struct {
	}
	defaultMutexValue struct {
		Expiry time.Time
	}
)

func (driver *defaultMutexDriver) Connect(name string, config chef.MutexConfig) (chef.MutexConnect, error) {
	setting := defaultMutexSetting{}
	return &defaultMutexConnect{
		name: name, config: config, setting: setting,
	}, nil
}

//打开连接
// 待处理，需要一个定时器，定期清理过期的数据
func (connect *defaultMutexConnect) Open() error {
	return nil
}

//关闭连接
func (connect *defaultMutexConnect) Close() error {
	return nil
}

//待优化，加上超时设置
func (connect *defaultMutexConnect) Lock(key string, expiry time.Duration) error {
	now := time.Now()

	if vv, ok := connect.locks.Load(key); ok {
		if tm, ok := vv.(defaultMutexValue); ok {
			if tm.Expiry.UnixNano() > now.UnixNano() {
				return errors.New("existed")
			}
		}
	}

	if expiry <= 0 {
		expiry = connect.config.Expiry
	}

	value := defaultMutexValue{
		Expiry: now.Add(connect.config.Expiry),
	}

	connect.locks.Store(key, value)

	return nil
}
func (connect *defaultMutexConnect) Unlock(key string) error {
	connect.locks.Delete(key)
	return nil
}
