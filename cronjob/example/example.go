package main

import (
	"fmt"
	"time"

	"github.com/better-go/pkg/cronjob"
)

func main() {
	// async way1:
	runAsyncWay1()

	// async way2:
	runAsyncWay2()

	// sync:
	//runSync()
}

func runAsyncWay1() {
	cj := cronjob.New()

	// register + run:
	cj.Run(
		cronjob.Task{
			Name:     "Test 1",
			Schedule: "@every 1s",
			TaskFunc: func() {
				fmt.Printf("test1, every 1s, %v\n", time.Now())
			},
		},
		cronjob.Task{
			Name:     "Test 2",
			Schedule: "@every 2s",
			TaskFunc: func() {
				fmt.Printf("\ttest2, every 2s, %v\n", time.Now())
			},
		},
	)

	// wait call cron task:
	time.Sleep(time.Minute * 5)
}

func runAsyncWay2() {
	cj := cronjob.New()

	// register:
	err := cj.RegisterTask(
		cronjob.Task{
			Name:     "Test 1",
			Schedule: "@every 1s",
			TaskFunc: func() {
				fmt.Printf("test1, every 1s, %v\n", time.Now())
			},
		},
		cronjob.Task{
			Name:     "Test 2",
			Schedule: "@every 2s",
			TaskFunc: func() {
				fmt.Printf("\ttest2, every 2s, %v\n", time.Now())
			},
		},
	)
	fmt.Printf("register err: %v\n", err)

	// run:
	cj.Run()

	// wait cron task:
	time.Sleep(time.Minute * 5)
}

func runSync() {
	cj := cronjob.New()

	// register:
	err := cj.RegisterTask(
		cronjob.Task{
			Name:     "Test 1",
			Schedule: "@every 1s",
			TaskFunc: func() {
				fmt.Printf("test1, every 1s, %v\n", time.Now())
			},
		},
		cronjob.Task{
			Name:     "Test 2",
			Schedule: "@every 2s",
			TaskFunc: func() {
				fmt.Printf("\ttest2, every 2s, %v\n", time.Now())
			},
		},
	)
	fmt.Printf("register err: %v\n", err)

	// run:
	cj.RunSync()

	// no need this
	//time.Sleep(time.Minute * 5)
}
