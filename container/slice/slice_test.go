package slice

import (
	"fmt"
	"testing"
)

func TestWithSlice(t *testing.T) {
	args := make([]interface{}, 0, 0)
	for i := 0; i < 50; i++ {
		args = append(args, fmt.Sprintf("arg-%d", i))
	}

	// task:
	tr := &TaskResource{
		Resource:  args,
		SliceSize: 3,
		Sleep:     100,
	}

	// fn:
	taskFn := func(slice []interface{}) error {
		t.Logf("slice args: %+v\n", slice)
		for _, item := range slice {
			v, ok := item.(string)
			if ok {
				t.Logf("\titem: %v\n", v)

			}
		}
		return nil
	}

	t.Log("slice ", WithSlice(tr, taskFn))
}
