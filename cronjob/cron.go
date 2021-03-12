package cronjob

import (
	"github.com/better-go/pkg/log"
	"github.com/robfig/cron/v3"
)

/*
cron job:
	- cron 表达式:
		- http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/tutorial-lesson-06.html
		- https://en.wikipedia.org/wiki/Cron
*/
type CronJob struct {
	Cron *cron.Cron // 暴露原库 API
}

func New() *CronJob {
	return &CronJob{
		Cron: cron.New(
			cron.WithSeconds(), // 支持解析秒
			cron.WithChain(),
		),
	}
}

// 注册 task:
func (m *CronJob) RegisterTask(tasks ...Task) (err error) {
	// batch register:
	for _, item := range tasks {
		// register:
		if entryID, err := m.Cron.AddFunc(item.Schedule, item.TaskFunc); err != nil {
			log.Errorf("cron job register tasks func error:, entryID=%v, err=%v", entryID, err)
		}
	}
	return err
}

// 注册和启动分开, 灵活调用位置
func (m *CronJob) Run(tasks ...Task) {
	m.RunAsync(tasks...)
}

// 异步:
func (m *CronJob) RunAsync(tasks ...Task) {
	// 允许在 run 中注册, 也可以分开, 传空即可
	_ = m.RegisterTask(tasks...)

	// 启动: 异步方式
	m.Cron.Start()
}

// 同步:
func (m *CronJob) RunSync(tasks ...Task) {
	// 允许在 run 中注册, 也可以分开, 传空即可
	_ = m.RegisterTask(tasks...)

	// 启动: 同步方式
	m.Cron.Run()
}

// stop:
func (m *CronJob) Stop() {
	m.Cron.Stop()
}
