package time

import (
	"time"
)

const (
	std = time.RFC3339 // 标准库格式

	// 时间格式: 参考标准库 (固定格式值, 不可乱改)
	fmtMonth = "2006-01"    // 月份格式
	fmtDay   = "2006-01-02" // 天格式
)
