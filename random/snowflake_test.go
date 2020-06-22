package random

import (
	"testing"
)

func TestSnowFlakeID(t *testing.T) {
	// batch:
	for i := 0; i < 100; i++ {
		uid := SnowFlakeID()
		t.Logf("gen unique id: %v", uid)
	}
}
