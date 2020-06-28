package retry

import (
	"errors"
	"fmt"
	"testing"
)

func TestWithRetry(t *testing.T) {

	// task func
	taskFn := func(i int) error {
		if i == 1 {
			fmt.Println("err, try again")
			return errors.New("err")
		}
		fmt.Println("no err")
		return nil
	}

	t.Log(WithRetry(2, DefaultBackOffConfig, func() error {
		return taskFn(1)
	}))

	t.Log(WithRetry(3, DefaultBackOffConfig, func() error {
		return taskFn(2)
	}))

}
