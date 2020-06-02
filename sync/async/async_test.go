package async

import (
	"context"
	"testing"
	"time"
)

func TestTaskDispatcher_Dispatch(t *testing.T) {
	ca := New("cache", Worker(1), Buffer(1024))
	var run bool

	// async task:
	ca.Dispatch(context.Background(), func(c context.Context) {
		run = true
		panic("error")
	})

	// sleep:
	time.Sleep(time.Millisecond * 50)
	t.Log("not panic")
	if !run {
		t.Fatal("expect run be true")
	}
}

func TestTaskDispatcher_Close(t *testing.T) {
	ca := New("cache", Worker(1), Buffer(1024))
	ca.Close()
	err := ca.Dispatch(context.Background(), func(c context.Context) {})
	if err == nil {
		t.Fatal("expect get err")
	}
}
