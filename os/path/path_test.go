package path

import "testing"

func TestCurrentPath(t *testing.T) {
	t.Log(RuntimePath())
	t.Log(CurrentDir())
}
