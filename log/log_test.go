package log

import "testing"

func TestInfo(t *testing.T) {
	Info("hello")
	Warnf("hello %v", "jim")
}

func TestError(t *testing.T) {
	Errorw("hello",
		"key:", "k1",
		"value:", "v1",
	)
}
