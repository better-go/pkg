package cronjob

// 一批任务:
type Tasks struct {
	Task []TaskItem
}

// 单个任务:
type TaskItem struct {
	Name     string
	Schedule string // 执行计划周期: cron 表达式 //
	TaskFunc func() // 任务方法
}
