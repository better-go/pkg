package cronjob

// 一批任务:
type Tasks struct {
	Task []Task
}

// 单个任务:
type Task struct {
	Name     string
	Schedule string // 执行计划周期: cron 表达式 //
	TaskFunc func() // 任务方法
}
