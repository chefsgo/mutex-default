package mutex

import (
	"github.com/chefsgo/chef"
)

func Driver() chef.MutexDriver {
	return &defaultMutexDriver{}
}

func init() {
	chef.Register("default", Driver())
}
