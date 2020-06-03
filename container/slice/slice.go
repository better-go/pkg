package slice

import (
	"time"
)

/*
分片执行:

*/

type TaskResource struct {
	Resource  []interface{} // 资源
	SliceSize int           // 切片大小
	Sleep     time.Duration // 延迟 N 毫秒执行[1s = 1000毫秒]
}

// 依赖闭包:
type TaskFunc func(slice []interface{}) error

func WithSlice(tr *TaskResource, taskFn TaskFunc) error {
	slice := make([]interface{}, 0, 0)

	// by slice:
	for _, item := range tr.Resource {
		// add:
		slice = append(slice, item)

		// do task:
		if len(slice) == tr.SliceSize {
			// run:
			_ = taskFn(slice)

			// clear:
			slice = make([]interface{}, 0, 0)

			//sleep:
			time.Sleep(tr.Sleep) // default = 0
		}
	}

	// handle last slice:
	if len(slice) > 0 {
		// run:
		_ = taskFn(slice)
	}

	return nil
}
