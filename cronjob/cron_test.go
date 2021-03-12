package cronjob

import (
	"fmt"
	"testing"
	"time"
)

// please check example usage: [this test not work)
func TestCronJob_Run(t *testing.T) {
	cj := New()

	// register + run:
	cj.Run(
		Task{
			Name:     "Test 1",
			Schedule: "@every 1s",
			TaskFunc: func() {
				fmt.Printf("test1, every 1s, %v", time.Now())
			},
		},
		Task{
			Name:     "Test 2",
			Schedule: "@every 2s",
			TaskFunc: func() {
				fmt.Printf("test2, every 2s, %v", time.Now())
			},
		},
	)

	// wait cron task:
	time.Sleep(time.Minute * 5)
}

func TestCronJob_RegisterTask(t *testing.T) {
	cj := New()

	err := cj.RegisterTask(Task{
		Name:     "Test 1",
		Schedule: "@every 1s",
		TaskFunc: func() {
			fmt.Printf("hello, every 1s, %v", time.Now())
		},
	})
	t.Logf("register err: %v", err)

	cj.Run()

	// wait cron task:
	time.Sleep(time.Minute * 5)
}
