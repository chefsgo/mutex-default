package mutex

import (
	"github.com/chefsgo/mutex"
)

func Driver() mutex.Driver {
	return &defaultDriver{}
}

func init() {
	mutex.Register("default", Driver())
}
