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
	cronJob *cron.Cron
}

func New() *CronJob {
	return &CronJob{
		cronJob: cron.New(
			cron.WithSeconds(), // 支持解析秒
			cron.WithChain(),
		),
	}
}

func (m *CronJob) RegisterTask(tasks ...TaskItem) (err error) {
	// batch register:
	for _, item := range tasks {
		// register:
		if entryID, err := m.cronJob.AddFunc(item.Schedule, item.TaskFunc); err != nil {
			log.Errorf("cron job register tasks func error:, entryID=%v, err=%v", entryID, err)
		}
	}
	return err
}

func (m *CronJob) Run(firstFn FirstRunFunc, scheduleFn ScheduleRunFunc) {

	go firstFn()

	//
	scheduleFn()
}

func (m *CronJob) runSchedule() {

}

type FirstRunFunc func()

type ScheduleRunFunc func()
