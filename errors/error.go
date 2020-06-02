package errors

import (
	"pkg/log"
)

// 重大 error, 需要主动 panic 的:
func PanicError(err error) {
	if err != nil {
		log.Errorf("error panic here, %+v", err)
		panic(err)
	}
	return
}
